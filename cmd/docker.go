package cmd

import (
	"sv-cli/metrics"
	"sv-cli/utils"

	"github.com/spf13/cobra"
)

func NewDockerCmd() *cobra.Command {
	dockerCmd := &cobra.Command{
		Use:   "docker",
		Short: "Manage Docker containers",
	}

	// Add subcommands
	dockerCmd.AddCommand(NewDockerShowCmd(), NewDockerListCmd())

	return dockerCmd
}

func NewDockerShowCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "show [container_name]",
		Short: "Check if a Docker container is running",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetDockerContainerInfo(args[0]), format)
		},
	}
}

func NewDockerListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all running Docker containers",
		Run: func(cmd *cobra.Command, args []string) {
			// Retrieve the format flag
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetDockerContainers(), format)
		},
	}
}
