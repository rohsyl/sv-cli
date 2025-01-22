package db

import (
	"github.com/spf13/cobra"
	"sv-cli/metrics"
	"sv-cli/utils"
)

func NewPostgresCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "postgres",
		Short: "Check PostgreSQL database connection",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			dsn, _ := cmd.Flags().GetString("dsn")
			utils.OutputResult(metrics.CheckDatabaseStatus("postgres", dsn), format)
		},
	}
}
