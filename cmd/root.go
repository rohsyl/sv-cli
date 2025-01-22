package cmd

import (
	"github.com/spf13/cobra"
	"sv-cli/cmd/db"
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
		NewCPUCmd(),
		NewRamCmd(),
		NewDiskCmd(),
		NewServiceCmd(),
		NewDockerCmd(),
		db.NewDbCmd(),
		NewSystemCmd(),
		NewSendMetricsCmd(),
	)

	return rootCmd
}
