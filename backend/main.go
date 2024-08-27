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
	appState := &api.AppState{DB: datab}
	api.Serve(appState)
}
