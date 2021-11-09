package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "matcher",
		Short: "A utility for verifying checking the result of input data against a match tree",
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
