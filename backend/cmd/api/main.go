package main

import (
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal"
	"github.com/jcocozza/poop.map/backend/internal/config"
	"github.com/jcocozza/poop.map/backend/internal/logger"
)



func main() {
	lggr := logger.CreateLogger(slog.LevelDebug, config.Prod)
	cfg := config.ReadConfig("../../../config.json", lggr)
	r, _ := internal.SetupAPI(lggr, cfg)
	lggr.Info("running on 8111")
	if err := http.ListenAndServe(":8111", r); err != nil {
		panic(err)
	}
}
