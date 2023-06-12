// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package tools

import (
	"fmt"
	"github.com/playbymail/olytools/parsers/report"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

var convertReportToJsonCmd = &cobra.Command{
	Use:   "report-to-json",
	Short: "convert report to JSON",
	Long:  `Convert report to JSON.`,
	Run: func(cmd *cobra.Command, args []string) {
		toolsCfg.Convert.Import = filepath.Clean(toolsCfg.Convert.Import)
		fmt.Printf("import %s as REPORT\n", toolsCfg.Convert.Import)
		toolsCfg.Convert.Export = filepath.Clean(toolsCfg.Convert.Export)
		fmt.Printf("export %s as JSON\n", toolsCfg.Convert.Export)
		if toolsCfg.Convert.Log != "" {
			toolsCfg.Convert.Log = filepath.Clean(toolsCfg.Convert.Log)
			fmt.Printf("log    %s\n", toolsCfg.Convert.Log)
		}

		input, err := os.ReadFile(toolsCfg.Convert.Import)
		if err != nil {
			log.Fatal(err)
		}
		parser, err := report.NewParser(input)
		if err != nil {
			log.Fatal(err)
		}
		if toolsCfg.Convert.Log != "" {
			w, err := os.OpenFile(toolsCfg.Convert.Log, os.O_CREATE, 0644)
			if err != nil {
				log.Fatal(err)
			}
			parser.Dump(w)
			if err := w.Close(); err != nil {
				log.Fatal(err)
			}
		}
		rpt, err := report.Parse(input)
		if err != nil {
			log.Fatal(err, "\n\n", rpt)
		} else if rpt == nil {
			log.Fatal("rpt is nil")
		}
	},
}
