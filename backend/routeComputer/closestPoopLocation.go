package routecomputer

import (
	"context"
	"fmt"

	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/model"
)

func computeClosestPoopLocation(currLat, currLong float64, poopLocations []model.PoopLocation) model.PoopLocation {
	smallestDist := 0.0

	var closestPoopLocation model.PoopLocation
	for i, poopLocation := range poopLocations {
		haversineDist := haversine(currLat, currLong, poopLocation.Latitude, poopLocation.Longitude)

		// why am i so lazy
		if i == 0 {
			smallestDist = haversineDist
			closestPoopLocation = poopLocation
		} else if haversineDist < smallestDist {
			closestPoopLocation = poopLocation
		}
	}
	return closestPoopLocation
}

func GetClosestPoopLocationAndRoute(ctx context.Context, currLat, currLong float64, repo *database.Repository) (model.PoopLocation, string, error) {
	poopLocations, err := repo.ListPoopLocations(ctx)
	if err != nil {
		return model.PoopLocation{}, "", err
	}

	closestPoopLocation := computeClosestPoopLocation(currLat, currLong, poopLocations)

	route, err := GetRoute(currLat, currLong, closestPoopLocation.Latitude, closestPoopLocation.Longitude)
	if err != nil {
		return closestPoopLocation, route, fmt.Errorf("error computing route. try again later. %w", err)
	}
	return closestPoopLocation, route, nil
}
