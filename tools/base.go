// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package tools

import (
	"github.com/spf13/cobra"
)

// baseCmd represents the base command when called without any subcommands
var baseCmd = &cobra.Command{
	Short:   "olytools",
	Long:    `Tools to help play the game of Olympia.`,
	Version: "0.0.1",
}
