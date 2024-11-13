package api

import (
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api/handlers"
	"github.com/jcocozza/poop.map/backend/internal/api/middleware"
	"github.com/rs/cors"
)

func NewRouter(logger *slog.Logger, apiKey string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.Status)

	// adding middlewares
	apiKeyProtected := middleware.AuthorizationMiddleware(mux, logger, apiKey)
	loggedMux := middleware.LoggingMiddleware(apiKeyProtected, logger)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: Limit this to only the web server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-User-UUID"},
	})
	final := c.Handler(loggedMux)
	return final
}
