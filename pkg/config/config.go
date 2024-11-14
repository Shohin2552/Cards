package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	PostgresURL string `json:"postgres_url"`
}

func NewConfig() *Config {
	var config *Config

	content, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	return config
}
