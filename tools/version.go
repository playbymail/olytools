// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package tools

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show application version",
	Long:  `Show application version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", toolsCfg.Version)
	},
}
