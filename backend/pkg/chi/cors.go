package chi

import (
	"github.com/go-chi/cors"
	"steller-api/pkg/config"
)

func NewCors(cfg *config.Config) cors.Options {
	return cors.Options{
		AllowedOrigins:   []string{"capacitor://localhost", "https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}
}
