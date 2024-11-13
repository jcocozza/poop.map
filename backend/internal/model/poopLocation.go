package model

import "time"

type LocationType struct {
	Name  string
	Color string
}

var (
	Regular    LocationType = LocationType{"regular", "red"}
	PortaPotty LocationType = LocationType{"porta potty", "green"}
	Outhouse   LocationType = LocationType{"outhouse", "brown"}
	Other      LocationType = LocationType{"other", "blue"}
)

type Season string

const (
	Summer Season = "summer"
	Fall   Season = "fall"
	Winter Season = "winter"
	Spring Season = "spring"
)

type PoopLocation struct {
	UUID         string       `json:"uuid"`
	Latitude     float64      `json:"latitude"`
	Longitude    float64      `json:"longitude"`
	FirstCreated time.Time    `json:"first_created"`
	LocationType LocationType `json:"location_type"`
	Name         string       `json:"name"`
	Seasonal     bool         `json:"seasonal"`
	Seasons      []Season     `json:"seasons"`
}
