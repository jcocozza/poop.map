package database

import (
	"context"

	"github.com/jcocozza/poop.map/backend/model"
)

type Repository struct {
	DB *Database
}

// Create a new poop location in the database
func (r *Repository) CreatePoopLocation(ctx context.Context, poopLocation model.PoopLocation) error {
	sql := `INSERT INTO poop_locations (uuid, latitude, longitude, rating, first_created, location_type, name, notes)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	return r.DB.Execute(ctx, sql, poopLocation.Uuid, poopLocation.Latitude, poopLocation.Longitude, poopLocation.Rating, poopLocation.FirstCreated, poopLocation.LocationType, poopLocation.Name, poopLocation.Notes)
}

// List all poop locations in the database
func (r *Repository) ListPoopLocations(ctx context.Context) ([]model.PoopLocation, error) {
	sql := "SELECT uuid, latitude, longitude, rating, first_created, location_type, name, notes FROM poop_locations;"
	poopLocations := []model.PoopLocation{}
	rows, err := r.DB.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		poopLocation := model.PoopLocation{}

		err := rows.Scan(&poopLocation.Uuid, &poopLocation.Latitude, &poopLocation.Longitude, &poopLocation.Rating, &poopLocation.FirstCreated, &poopLocation.LocationType, &poopLocation.Name, &poopLocation.Notes)
		if err != nil {
			continue
		}
		poopLocations = append(poopLocations, poopLocation)
	}
	return poopLocations, nil
}
