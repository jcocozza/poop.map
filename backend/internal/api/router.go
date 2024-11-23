package api

import (
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api/handlers"
	"github.com/jcocozza/poop.map/backend/internal/api/middleware"
	"github.com/rs/cors"
)

func NewRouter(
	logger *slog.Logger,
	apiKey string,
	plh *handlers.PoopLocationHandler,
	pluh *handlers.PoopLocationUUIDHandler,
	rh *handlers.ReviewHandler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", handlers.Status)

	mux.HandleFunc("/poop-location", plh.PoopLocationHandler)

	mux.HandleFunc("/poop-location/{uuid}", pluh.PoopLocationUUIDHandler)
	mux.HandleFunc("/poop-location/{uuid}/upvote", pluh.Upvote)
	mux.HandleFunc("/poop-location/{uuid}/downvote", pluh.Downvote)

	mux.HandleFunc("/poop-location/{uuid}/review", rh.ReviewHandler)

	mux.HandleFunc("/review/{uuid}/upvote", rh.Upvote)
	mux.HandleFunc("/review/{uuid}/downvote", rh.Downvote)

	// adding middlewares
	apiKeyProtected := middleware.AuthorizationMiddleware(mux, logger, apiKey)
	loggedMux := middleware.LoggingMiddleware(apiKeyProtected, logger)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // TODO: Limit this to only the web server
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})
	final := c.Handler(loggedMux)
	return final
}
