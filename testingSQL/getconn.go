package belajar_go_sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetConn() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gosql")
	if err != nil {
		panic(err)
	}
	return db
}
