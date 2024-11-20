package testing

import (
	"log/slog"
	"net/http/httptest"
	"os"

	"github.com/jcocozza/poop.map/backend/internal"
	"github.com/jcocozza/poop.map/backend/internal/config"
	"github.com/jcocozza/poop.map/backend/internal/logger"
)

func SetupTest() *httptest.Server {
	lggr := logger.CreateLogger(slog.LevelDebug, config.Test)
	r, db := internal.SetupAPI(lggr, config.Test)

	// read in test sql
	sqlbytes, err := os.ReadFile("test_locations.sql")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(sqlbytes))
	if err != nil {
		panic(err)
	}

	// read in test sql
	sqlbytes2, err := os.ReadFile("test_reviews.sql")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(sqlbytes2))
	if err != nil {
		panic(err)
	}
	return httptest.NewServer(r)
}
