package metrics

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func CheckDatabaseStatus(dbType, dsn string) MetricResult {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return MetricResult{Success: false, Error: err.Error()}
	}

	return MetricResult{Success: true, Data: DatabaseStatus{Type: dbType, Status: "up"}}
}
