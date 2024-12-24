
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB *sql.DB

func InitMySQL(dsn string) error {
	var err error
	mysqlDB, err = sql.Open("mysql", dsn)
	if (err != nil) {
		return err
	}
	return mysqlDB.Ping()
}

func GetMySQLDB() *sql.DB {
	return mysqlDB
}

func CloseMySQL() error {
	return mysqlDB.Close()
}