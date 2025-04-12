package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"sv-cli/cmd/db"
    "sv-cli/internal"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "sv",
		Short: "CLI tool for system information",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
	        versionFlag, _ := cmd.Flags().GetBool("version")
	        if versionFlag {
	            fmt.Println("sv-cli " + internal.Version)
	            os.Exit(0)
	        }
	    },
	}

	// Persistent flags
	rootCmd.PersistentFlags().String("format", "json", "Output format: json or table")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number")

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
		VersionCmd(),
	)

	return rootCmd
}
