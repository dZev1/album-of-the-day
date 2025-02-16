package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfig struct {
	ConnStr string `json:"connStr"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
