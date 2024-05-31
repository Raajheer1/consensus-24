package models

import (
	"log"
	"steller-api/pkg/database"
)

func AutoMigrate() {
	err := database.DB.AutoMigrate(
		Account{},
		Feed{},
		Loan{},
		Vote{},
		GovernanceVote{},
		LoanLender{},
	)
	if err != nil {
		log.Fatal("[Database] Migration Error:", err)
	}
}
