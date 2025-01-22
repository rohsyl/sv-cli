package cmd

import (
	"github.com/spf13/cobra"
	"sv-cli/metrics"
	"sv-cli/utils"
)

func NewCPUCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "cpu",
		Short: "Get CPU usage",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetCPUUsage(), format)
		},
	}
}
