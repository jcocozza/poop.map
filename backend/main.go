package main

import (
	"github.com/jcocozza/poop.map/backend/api"
	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/utils"
)

const cfgPath = "../config.json"

func main() {
	cfg := utils.ReadConfig(cfgPath)

	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	datab := &database.Database{DB: db}
	repo := &database.Repository{DB: datab}
	appState := &api.AppState{R: repo, Cfg: cfg}
	api.Serve(appState)
}
