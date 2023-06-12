// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package config creates a single data structure for the tool configuration.
package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/playbymail/olytools/semver"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env struct {
		Environment string
	}
	Convert struct {
		Type   string
		Import string
		Export string
		Log    string
	}
	Random struct {
		Seed int64
	}
	Version semver.Version
}

// Default returns a config with default values.
func Default() *Config {
	cfg := &Config{}
	cfg.Random.Seed = time.Now().UnixNano()
	cfg.Version = semver.New(0, 0, 1)
	return cfg
}

// New returns a configuration initialized with default values
// and updated from the environment.
func New() (*Config, error) {
	cfg := Default()

	cfg.Env.Environment = strings.TrimSpace(strings.ToLower(os.Getenv("OLYTOOL_ENV")))
	if cfg.Env.Environment == "" {
		cfg.Env.Environment = "dev"
	}
	if err := godotenv.Load(".env.local"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config: %w", err)
	}
	if err := godotenv.Load(fmt.Sprintf(".env.%s", cfg.Env.Environment)); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config: %w", err)
	}
	if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("config: %w", err)
	}

	return cfg, nil
}
