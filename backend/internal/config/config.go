package config

import (
	"fmt"
	"log/slog"
)

type Environment string

const (
	Test Environment = "test"
	Dev  Environment = "dev"
	Prod Environment = "prod"
)

const api_key = "test_api_key"
const base_db_path = "poop_locations_database_%s.db"

type Config struct {
	APIKey       string
	DatabasePath string
	Env          Environment
}

func NewConfig(env Environment, apiKey string) Config {
	var dbPath string
	switch env {
	case Test:
		dbPath = fmt.Sprintf(base_db_path, Test)
	case Dev:
		dbPath = fmt.Sprintf(base_db_path, Dev)
	case Prod:
		dbPath = fmt.Sprintf(base_db_path, Prod)
	}
	return Config{
		APIKey: apiKey,
		DatabasePath: dbPath,
		Env: env,
	}
}

func ReadConfig(logger *slog.Logger, env Environment) Config {
	cfg := NewConfig(env, api_key)
	logger.Info(fmt.Sprintf("Environment is: %s", cfg.Env))
	logger.Info(fmt.Sprintf("API KEY is: %s", cfg.APIKey))
	logger.Info(fmt.Sprintf("database path is: %s", cfg.DatabasePath))
	return cfg
}
