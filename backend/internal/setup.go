package internal

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/jcocozza/poop.map/backend/internal/api"
	"github.com/jcocozza/poop.map/backend/internal/api/handlers"
	"github.com/jcocozza/poop.map/backend/internal/config"
	"github.com/jcocozza/poop.map/backend/internal/database"
	"github.com/jcocozza/poop.map/backend/internal/repository/sqlite"
	"github.com/jcocozza/poop.map/backend/internal/service"
)

func SetupAPI(lggr *slog.Logger, env config.Environment) (http.Handler, *sql.DB) {
	cfg := config.ReadConfig(lggr, env)
	db := database.NewSQLiteDB(cfg.DatabasePath).Connect()
	plr := sqlite.NewSQLitePoopLocationRepository(db, lggr)
	rr := sqlite.NewSQLiteReviewRepository(db, lggr)
	pls := service.NewPoopLocationService(plr, lggr)
	rs := service.NewReviewService(rr, lggr)
	plh := handlers.NewPoopLocationHandler(pls, lggr)
	pluh := handlers.NewPoopLocationUUIDHandler(pls, lggr)
	rh := handlers.NewReviewHandler(rs, lggr)
	return api.NewRouter(lggr, cfg.APIKey, plh, pluh, rh), db
}
