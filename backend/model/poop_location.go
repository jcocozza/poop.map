package model

type PoopLocation struct {
	Uuid         string  `json:"uuid"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Rating       int     `json:"rating"`
	FirstCreated string  `json:"first_created"`
	LocationType string  `json:"location_type"`
	Name         string  `json:"name"`
	Notes		 string  `json:"notes"`
}
