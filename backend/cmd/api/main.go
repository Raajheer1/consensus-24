package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"steller-api/pkg/chi"
	"steller-api/pkg/config"
	"steller-api/pkg/database"
	"steller-api/pkg/database/models"
	"steller-api/pkg/scheduler"
	v1 "steller-api/v1"
	"syscall"
	"time"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize configuration
	config.Cfg = config.New()

	// Connect to the database
	database.DB = database.Connect(config.Cfg.Database)
	models.AutoMigrate()

	// Setup the router
	r := chi.New(config.Cfg)
	v1.Router(r)

	// Setup the HTTP server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3000),
		Handler: r,
	}

	// Run the HTTP server in a separate goroutine
	go func() {
		log.Printf("Starting HTTP server on port %d", 3000)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting HTTP server: %s", err)
		}
	}()

	// Initialize the task scheduler
	taskRunner := scheduler.NewScheduler()

	// Task to check and see if any governance votes have ended
	taskRunner.AddTask(&scheduler.Task{
		ID:   1,
		Name: "Governance Vote Check",
		TaskFunc: func() {
			activeGVs, err := models.GetActiveGVs()
			if err != nil {
				log.Printf("Error getting active governance votes: %v", err)
				return
			}

			for _, gv := range activeGVs {
				if err := gv.CompletionCheck(); err != nil {
					log.Printf("Error checking completion for governance vote #%d: %v", gv.ID, err)
				}
			}
		},
		Interval: time.Minute,
	})

	taskRunner.AddTask(&scheduler.Task{
		ID:   2,
		Name: "Loan Validator",
		TaskFunc: func() {
			loans, err := models.GetLoans()
			if err != nil {
				log.Printf("Error getting loans: %v", err)
				return
			}

			for _, loan := range loans {
				amountRaised := 0.0
				for _, loanLender := range loan.Lenders {
					amountRaised += loanLender.LoanAmount
				}

				if amountRaised != loan.AmountRaised {
					log.Printf("Amount raised for loan #%d does not match the sum of the loan lenders: %d != %d", loan.ID, loan.AmountRaised, amountRaised)
				}
			}
		},
		Interval: 10 * time.Minute,
	})

	// Setup signal catching for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Wait for an interrupt signal
	<-sigs

	// Graceful shutdown procedures
	fmt.Println("Received interrupt signal, shutting down...")

	// Shutdown the HTTP server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("HTTP server Shutdown: %v", err)
	}

	// Stop the task scheduler
	taskRunner.Stop()

	fmt.Println("All tasks stopped, exiting.")
}
