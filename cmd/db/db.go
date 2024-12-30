package db

import (
	"github.com/spf13/cobra"
)

func NewDbCmd() *cobra.Command {
	dbCmd := &cobra.Command{
		Use:   "db",
		Short: "Check database connectivity",
	}
	dbCmd.PersistentFlags().String("dsn", "", "Database connection string")
	dbCmd.MarkPersistentFlagRequired("dsn")
	dbCmd.AddCommand(NewMysqlCmd(), NewPostgresCmd(), NewMssqlCmd())
	return dbCmd
}
