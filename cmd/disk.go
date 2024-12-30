package cmd

import (
	"fmt"
	"sv-cli/utils"

	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

func NewDiskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Get disk usage",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			usage, _ := disk.Usage("/")

			if format == "json" {
				utils.PrintJSON(usage)
			} else {
				utils.PrintTable([][]string{
					{"Path", usage.Path},
					{"Total", fmt.Sprintf("%v GB", usage.Total/1024/1024/1024)},
					{"Used", fmt.Sprintf("%v GB", usage.Used/1024/1024/1024)},
					{"Free", fmt.Sprintf("%v GB", usage.Free/1024/1024/1024)},
					{"UsedPercent", fmt.Sprintf("%.2f%%", usage.UsedPercent)},
				}, []string{"Metric", "Value"})
			}
		},
	}
}
