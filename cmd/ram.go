package cmd

import (
	"fmt"
	"sv-cli/utils"

	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

func NewRamCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ram",
		Short: "Get RAM usage",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			v, _ := mem.VirtualMemory()

			if format == "json" {
				utils.PrintJSON(v)
			} else {
				utils.PrintTable([][]string{
					{"Total", fmt.Sprintf("%v MB", v.Total/1024/1024)},
					{"Used", fmt.Sprintf("%v MB", v.Used/1024/1024)},
					{"Free", fmt.Sprintf("%v MB", v.Free/1024/1024)},
					{"UsedPercent", fmt.Sprintf("%.2f%%", v.UsedPercent)},
				}, []string{"Metric", "Value"})
			}
		},
	}
}
