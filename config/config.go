package config

import (
	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-search-reindex-api
type Config struct {
	BindAddr string `envconfig:"BIND_ADDR"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr: "localhost:8080",
	}

	return cfg, envconfig.Process("", cfg)
}
