package cmd

import (
	"sv-cli/metrics"
	"sv-cli/utils"

	"github.com/spf13/cobra"
)

func NewDiskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Get disk usage",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetDiskUsage(), format)
		},
	}
}
