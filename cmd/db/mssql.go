package db

import (
	"database/sql"
	"sv-cli/utils"

	"github.com/spf13/cobra"
	_ "github.com/denisenkom/go-mssqldb"
)

func NewMssqlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mssql",
		Short: "Check MSSQL database connection",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			dsn, _ := cmd.Flags().GetString("dsn")
			result := map[string]string{"database": "MSSQL"}

			db, err := sql.Open("sqlserver", dsn)
			if err != nil {
				result["status"] = "failed"
				result["message"] = err.Error()
			} else if err := db.Ping(); err != nil {
				result["status"] = "failed"
				result["message"] = err.Error()
			} else {
				result["status"] = "success"
				result["message"] = "Connection successful"
			}
			db.Close()

			if format == "json" {
				utils.PrintJSON(result)
			} else {
				utils.PrintTable([][]string{{result["database"], result["status"], result["message"]}}, []string{"Database", "Status", "Message"})
			}
		},
	}
}
