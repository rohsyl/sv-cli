package cmd

import (
	"sv-cli/metrics"
	"sv-cli/utils"

	"github.com/spf13/cobra"
)

func NewServiceCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "service [name]",
		Short: "Check if a service is running",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetServiceStatus(args[0]), format)
		},
	}
}
