package api

import (
	"fmt"
	"net/http"
)

func Serve(appState *AppState) {
	http.HandleFunc("/api/create", appState.createPoopLocation)
	http.HandleFunc("/api/list_all", appState.listPoopLocations)

	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
