package belajar_go_sql

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestGetConn(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gosql")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
