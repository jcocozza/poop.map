package model

import "time"

type Review struct {
	UUID             string    `json:"uuid"`
	PoopLocationUUID string    `json:"poop_location_uuid"`
	Rating           int       `json:"rating"`
	Comment          string    `json:"comment"`
	Time             time.Time `json:"time"`
	Upvotes          int       `json:"upvotes"`
	DownVotes        int       `json:"downvotes"`
}
