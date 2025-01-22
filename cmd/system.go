package cmd

import (
	"github.com/spf13/cobra"
	"sv-cli/metrics"
	"sv-cli/utils"
)

func NewSystemCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "system",
		Short: "Get system information",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetSystemInfo(), format)
		},
	}
}
