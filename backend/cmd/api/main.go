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
	r, _ := internal.SetupAPI(lggr, config.Prod)
	lggr.Info("running on 8111")
	if err := http.ListenAndServe(":8111", r); err != nil {
		panic(err)
	}
}
