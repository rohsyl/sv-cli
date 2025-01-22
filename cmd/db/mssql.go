package db

import (
	"sv-cli/metrics"
	"sv-cli/utils"

	"github.com/spf13/cobra"
)

func NewMssqlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mssql",
		Short: "Check MSSQL database connection",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			dsn, _ := cmd.Flags().GetString("dsn")
			utils.OutputResult(metrics.CheckDatabaseStatus("postgres", dsn), format)
		},
	}
}
