package config

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

const base_db_path = "poop_locations_database_%s.db"

type Environment string
const (
	Test Environment = "test"
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

func validateEnvironment(env Environment) {
	if env != Test && env != Dev && env != Prod {
		panic("invalid environment setting")
	}
}

type Config struct {
	APIKey string      `json:"api_key"`
	Env    Environment `json:"environment"`

	DatabasePath string
}

func createDBPath(env Environment) string {
	var dbPath string
	switch env {
	case Test:
		dbPath = fmt.Sprintf(base_db_path, Test)
	case Dev:
		dbPath = fmt.Sprintf(base_db_path, Dev)
	case Prod:
		dbPath = fmt.Sprintf(base_db_path, Prod)
	}
	return dbPath
}

func ReadConfig(path string, logger *slog.Logger) Config {
	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var cfg Config
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		panic(err)
	}
	validateEnvironment(cfg.Env)
	cfg.DatabasePath = createDBPath(cfg.Env)
	logger.Info(fmt.Sprintf("environment is: %s", cfg.Env))
	logger.Info(fmt.Sprintf("api key is: %s", cfg.APIKey))
	logger.Info(fmt.Sprintf("database path is: %s", cfg.DatabasePath))
	return cfg
}
