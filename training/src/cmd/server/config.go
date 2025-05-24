package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port string `env:"PORT"`
}

func loadCfgFromEnv() (Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return Config{}, fmt.Errorf("parse environment: %w", err)
	}
	return cfg, nil
}
