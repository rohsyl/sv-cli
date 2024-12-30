package cmd

import (
	"os/exec"
	"runtime"
	"strings"
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
			service := args[0]

			var status string
			if runtime.GOOS == "linux" {
				out, _ := exec.Command("systemctl", "is-active", service).Output()
				status = strings.TrimSpace(string(out))
			} else if runtime.GOOS == "windows" {
				out, _ := exec.Command("sc", "query", service).Output()
				if strings.Contains(string(out), "RUNNING") {
					status = "running"
				} else {
					status = "not running"
				}
			} else {
				status = "unsupported"
			}

			result := map[string]string{"service": service, "status": status}
			if format == "json" {
				utils.PrintJSON(result)
			} else {
				utils.PrintTable([][]string{{service, status}}, []string{"Service", "Status"})
			}
		},
	}
}
