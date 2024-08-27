package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/model"
)

type AppState struct {
	DB *database.Database
}

func (a *AppState) createPoopLocation(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	sql := `INSERT INTO poop_locations (uuid, latitude, longitude, rating, first_created, location_type, name)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	var poopLocation model.PoopLocation
	err := json.NewDecoder(r.Body).Decode(&poopLocation)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = a.DB.Execute(context.TODO(), sql, poopLocation.Uuid, poopLocation.Latitude, poopLocation.Longitude, poopLocation.Rating, poopLocation.FirstCreated, poopLocation.LocationType, poopLocation.Name)
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

	sql := "SELECT uuid, latitude, longitude, rating, first_created, location_type, name FROM poop_locations;"

	rows, err := a.DB.Query(context.TODO(), sql)
	if err != nil {
		http.Error(w, "failed to query database", http.StatusInternalServerError)
		return
	}

	poopLocations := []model.PoopLocation{}
	for rows.Next() {
		poopLocation := model.PoopLocation{}

		err := rows.Scan(&poopLocation.Uuid, &poopLocation.Latitude, &poopLocation.Longitude, &poopLocation.Rating, &poopLocation.FirstCreated, &poopLocation.LocationType, &poopLocation.Name)
		if err != nil {
			continue
		}
		poopLocations = append(poopLocations, poopLocation)
	}

	err = json.NewEncoder(w).Encode(poopLocations)
	if err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
