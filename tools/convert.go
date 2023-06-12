// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package tools

import (
	"fmt"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert application data",
	Long:  `Convert application data.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", toolsCfg.Version)
	},
}
