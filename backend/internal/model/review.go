package model

import "time"

type Review struct {
	UUID             string    `json:"uuid"`
	PoopLocationUuid string    `json:"poop_location_uuid"`
	Rating           int       `json:"rating"`
	Comment          string    `json:"comment"`
	Time             time.Time `json:"time"`
}
