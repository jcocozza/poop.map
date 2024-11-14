package model

import "time"

type LocationType string

const (
	Regular    LocationType = "regular"
	PortaPotty LocationType = "porta potty"
	Outhouse   LocationType = "outhouse"
	Other      LocationType = "other"
)

// we can use a single binary number to encode the seasons
const (
	summer = 1 << iota
	fall
	winter
	spring
)

type Season string

const (
	Summer Season = "summer"
	Fall   Season = "fall"
	Winter Season = "winter"
	Spring Season = "spring"
)

// map the number (in binary) to a list of seasons
func GetSeasons(mask int) []Season {
	var seasons []Season
	if mask&summer != 0 {
		seasons = append(seasons, Summer)
	}
	if mask&fall != 0 {
		seasons = append(seasons, Fall)
	}
	if mask&winter != 0 {
		seasons = append(seasons, Winter)
	}
	if mask&spring != 0 {
		seasons = append(seasons, Spring)
	}
	return seasons
}

func SeasonMask(seasons []Season) int {
	total := 0
	for _, s := range seasons {
		switch s {
		case Summer:
			total += summer
		case Fall:
			total += fall
		case Winter:
			total += winter
		case Spring:
			total += spring
		}
	}
	return total
}

type PoopLocation struct {
	UUID         string       `json:"uuid"`
	Name         string       `json:"name"`
	Latitude     float64      `json:"latitude"`
	Longitude    float64      `json:"longitude"`
	FirstCreated time.Time    `json:"first_created"`
	LastModified time.Time    `json:"last_modified"`
	LocationType LocationType `json:"location_type"`
	Seasonal     bool         `json:"seasonal"`
	Seasons      []Season     `json:"seasons"`
	Accessible   bool         `json:"accessible"`
	Upvotes      int          `json:"upvotes"`
	DownVotes    int          `json:"downvotes"`
}
