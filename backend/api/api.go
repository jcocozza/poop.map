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

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000"}, // Allow your frontend origin
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
	})

	handler := c.Handler(mux)

	/*
	http.HandleFunc("/api/create", appState.createPoopLocation)
	http.HandleFunc("/api/list_all", appState.listPoopLocations)
	*/

	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
