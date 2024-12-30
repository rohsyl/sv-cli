package db

import (
	"database/sql"
	"sv-cli/utils"

	"github.com/spf13/cobra"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "mysql",
		Short: "Check MySQL database connection",
		Run: func(cmd *cobra.Command, args []string) {
			format, _ := cmd.Flags().GetString("format")
			dsn, _ := cmd.Flags().GetString("dsn")
			result := map[string]string{"database": "MySQL"}

			db, err := sql.Open("mysql", dsn)
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
