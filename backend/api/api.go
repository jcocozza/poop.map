package api

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
)

func Serve(appState *AppState) {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/list_all", appState.listPoopLocations)
	mux.HandleFunc("/api/create", appState.createPoopLocation)
	mux.HandleFunc("/api/closest", appState.computeClosestPoopLocation)

	frontendOrigin := fmt.Sprintf("%s", appState.Cfg.FrontendUrl)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{frontendOrigin}, // Allow your frontend origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	port := fmt.Sprintf(":%s", appState.Cfg.BackendPort)
	appState.Logger.Info(fmt.Sprintf("Starting server on port %s...\n", port))
	if err := http.ListenAndServe(port, handler); err != nil {
		appState.Logger.Error(fmt.Sprintf("Server failed to start: %v\n", err))
	}
}
