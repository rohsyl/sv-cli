package db

import (
	"sv-cli/metrics"
	"sv-cli/utils"

	"github.com/spf13/cobra"
)

func NewMysqlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mysql",
		Short: "Check MySQL database connection",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			dsn, _ := cmd.Flags().GetString("dsn")
			utils.OutputResult(metrics.CheckDatabaseStatus("mysql", dsn), format)
		},
	}
}
