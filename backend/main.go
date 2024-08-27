package main

import (
	"github.com/jcocozza/poop.map/backend/api"
	"github.com/jcocozza/poop.map/backend/database"
)

func main() {
	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	datab := &database.Database{DB: db}
	repo := &database.Repository{DB: datab}
	appState := &api.AppState{R: repo}
	api.Serve(appState)
}
