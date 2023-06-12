// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package tools implements a Cobra CLI for olytools.
package tools

import (
	"fmt"
	"github.com/playbymail/olytools/config"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the base Command.
func Execute(cfg *config.Config) error {
	toolsCfg = cfg

	baseCmd.AddCommand(convertCmd)
	convertReportToJsonCmd.Flags().StringVar(&toolsCfg.Convert.Export, "export", "", "file to export")
	if err := convertReportToJsonCmd.MarkFlagRequired("export"); err != nil {
		return fmt.Errorf("cobra: init: report-to-json: %w", err)
	}
	convertReportToJsonCmd.Flags().StringVar(&toolsCfg.Convert.Import, "import", "", "file to import")
	if err := convertReportToJsonCmd.MarkFlagRequired("import"); err != nil {
		return fmt.Errorf("cobra: init: report-to-json: %w", err)
	}
	convertReportToJsonCmd.Flags().StringVar(&toolsCfg.Convert.Log, "log", "", "logging file (optional)")
	convertCmd.AddCommand(convertReportToJsonCmd)

	baseCmd.AddCommand(versionCmd)

	return baseCmd.Execute()
}

// toolsCfg is set by the Execute() function.
var toolsCfg *config.Config
