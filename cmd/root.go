package cmd

import (
	"sv-cli/cmd/db"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sv",
		Short: "CLI tool for system information",
	}

	// Persistent flags
	rootCmd.PersistentFlags().String("format", "table", "Output format: json or table")

	// Add subcommands
	rootCmd.AddCommand(
		NewRamCmd(),
		NewDiskCmd(),
		NewServiceCmd(),
		NewDockerCmd(),
		db.NewDbCmd(),
	)

	return rootCmd
}
