package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/model"
	routecomputer "github.com/jcocozza/poop.map/backend/routeComputer"
	"github.com/jcocozza/poop.map/backend/utils"
)

type AppState struct {
	Logger *slog.Logger
	R   *database.Repository
	Cfg utils.Config
}

// a requestion to this handle looks like:
// http://localhost:8080/location?latitude=40.7128&longitude=-74.0060
func (a *AppState) computeClosestPoopLocation(w http.ResponseWriter, r *http.Request) {
	a.Logger.Debug("computing closest poop location")
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	latitudeStr := r.URL.Query().Get("latitude")
	longitudeStr := r.URL.Query().Get("longitude")
	a.Logger.Debug("", slog.String("latitude", latitudeStr), slog.String("longitude", longitudeStr))
	latitude, err := strconv.ParseFloat(latitudeStr, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid latitude", http.StatusBadRequest)
		return
	}
	longitude, err := strconv.ParseFloat(longitudeStr, 64)
	if err != nil {
		sendErrorResponse(w, "Invalid longitude", http.StatusBadRequest)
		return
	}

	poopLocation, route, err := routecomputer.GetClosestPoopLocationAndRoute(context.Background(), latitude, longitude, a.R)
	if err != nil {
		sendErrorResponse(w, "error getting closest poop location", http.StatusInternalServerError)
		return
	}
	type resp struct {
		PoopLocation model.PoopLocation `json:"poop_location"`
		Route        string             `json:"route"`
	}
	res := resp{PoopLocation: poopLocation, Route: route}
	a.Logger.Debug("preparing to return response", slog.Any("response", res))
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		sendErrorResponse(w, "Failed to encode", http.StatusInternalServerError)
		return
	}
}

func (a *AppState) createPoopLocation(w http.ResponseWriter, r *http.Request) {
	a.Logger.Debug("creating a poop location")
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var poopLocation model.PoopLocation
	err := json.NewDecoder(r.Body).Decode(&poopLocation)
	if err != nil {
		sendErrorResponse(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	a.Logger.Debug("recieved new poop location from sender", slog.Any("poop location", poopLocation))
	err = a.R.CreatePoopLocation(context.TODO(), poopLocation)
	if err != nil {
		sendErrorResponse(w, "Failed to write to database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AppState) listPoopLocations(w http.ResponseWriter, r *http.Request) {
	a.Logger.Debug("listing poop locations")
	if r.Method != http.MethodGet {
		sendErrorResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	poopLocations, err := a.R.ListPoopLocations(context.TODO())
	if err != nil {
		sendErrorResponse(w, "failed to query database", http.StatusInternalServerError)
		return
	}

	maxPoopLocations := 9
	if len(poopLocations) < maxPoopLocations {
		maxPoopLocations = len(poopLocations)
	}
	a.Logger.Debug("first 10 poop locations", slog.Any("first 10", poopLocations[:maxPoopLocations]))
	err = json.NewEncoder(w).Encode(poopLocations)
	if err != nil {
		sendErrorResponse(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}

// sendErrorResponse sends a JSON error response with the given message and status code.
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
