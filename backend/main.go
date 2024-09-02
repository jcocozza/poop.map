package main

import (
	"log/slog"
	"os"

	"github.com/jcocozza/poop.map/backend/api"
	"github.com/jcocozza/poop.map/backend/database"
	"github.com/jcocozza/poop.map/backend/utils"
)

const cfgPath = "../config.json"

func main() {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(handler)

	cfg := utils.ReadConfig(cfgPath)

	db, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}
	datab := &database.Database{DB: db, Logger: logger}
	repo := &database.Repository{DB: datab, Logger: logger}
	appState := &api.AppState{R: repo, Cfg: cfg, Logger: logger}
	api.Serve(appState)
}
