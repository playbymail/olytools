// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package main implements tools to help play Olympia.
package main

import (
	"github.com/playbymail/olytools/config"
	"github.com/playbymail/olytools/tools"
	"log"
	"math/rand"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(cfg.Random.Seed) // deprecated: use NewRand(NewSource(seed))

	if err := tools.Execute(cfg); err != nil {
		log.Fatal(err)
	}
}
