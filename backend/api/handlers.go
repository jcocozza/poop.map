package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/model"
	routecomputer "github.com/jcocozza/poop.map/backend/routeComputer"
)

type AppState struct {
	R *database.Repository
}

// a requestion to this handle looks like:
// http://localhost:8080/location?latitude=40.7128&longitude=-74.0060
func (a *AppState) computeClosestPoopLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	latitudeStr := r.URL.Query().Get("latitude")
	longitudeStr := r.URL.Query().Get("longitude")
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid latitude", http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		http.Error(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	poopLocation, route, err := routecomputer.GetClosestPoopLocationAndRoute(context.Background(), latitude, longitude, a.R)
	if err != nil {
		http.Error(w, "error getting closest poop location", http.StatusInternalServerError)
	}

	type resp struct {
		PoopLocation model.PoopLocation `json:"poop_location"`
		Route string `json:"route"`
	}

	res := resp{PoopLocation: poopLocation, Route: route}
	fmt.Println(res)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (a *AppState) createPoopLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var poopLocation model.PoopLocation
	err := json.NewDecoder(r.Body).Decode(&poopLocation)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = a.R.CreatePoopLocation(context.TODO(), poopLocation)
	if err != nil {
		http.Error(w, "Failed to write to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AppState) listPoopLocations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	poopLocations, err := a.R.ListPoopLocations(context.TODO())
	if err != nil {
		http.Error(w, "failed to query database", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(poopLocations)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
