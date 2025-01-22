package cmd

import (
	"github.com/spf13/cobra"
	"sv-cli/metrics"
	"sv-cli/utils"
)

func NewRamCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ram",
		Short: "Get RAM usage",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			utils.OutputResult(metrics.GetRAMUsage(), format)
		},
	}
}
