package config

import (
	"fmt"
	"log/slog"
)

const api_key = "test_api_key"
const db_path = "poop_locations_database.db"

type Config struct {
	APIKey       string
	DatabasePath string
}

func ReadConfig(logger *slog.Logger) Config {
	logger.Info(fmt.Sprintf("API KEY is: %s", api_key))
	logger.Info(fmt.Sprintf("database path is: %s", db_path))
	return Config{
		APIKey:       api_key,
		DatabasePath: db_path,
	}
}
