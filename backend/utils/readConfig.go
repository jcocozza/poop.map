package utils

import (
	"encoding/json"
	"os"
)

type Config struct {
	BackendUrl   string `json:"backend_url"`
	BackendPort  string `json:"backend_port"`
	FrontendUrl  string `json:"frontend_url"`
	FrontendPort string `json:"frontend_port"`
}

func ReadConfig(path string) Config {
	cfgBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err) // the config should always be there
	}
	var cfg Config
	err = json.Unmarshal(cfgBytes, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
