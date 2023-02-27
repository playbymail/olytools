// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package tools implements a Cobra CLI for olytools.
package tools

import (
	"github.com/playbymail/olytools/config"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the base Command.
func Execute(cfg *config.Config) error {
	toolsCfg = cfg

	baseCmd.AddCommand(versionCmd)

	return baseCmd.Execute()
}

// toolsCfg is set by the Execute() function.
var toolsCfg *config.Config
