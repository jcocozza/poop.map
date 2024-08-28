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

	frontendOrigin := fmt.Sprintf("%s:%s", appState.Cfg.FrontendUrl, appState.Cfg.FrontendPort)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{frontendOrigin}, // Allow your frontend origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	port := fmt.Sprintf(":%s", appState.Cfg.BackendPort)
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
