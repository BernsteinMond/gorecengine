package main

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	InferenceServer HTTPServer `envPrefix:"INFERENCE_SERVER_"`
}

type HTTPServer struct {
	ListenAddr string `env:"LISTEN_ADDR"`
}

func loadCfgFromEnv() (Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return Config{}, fmt.Errorf("parse environment: %w", err)
	}
	return cfg, nil
}
