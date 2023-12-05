package belajar_go_sql

import (
	"fmt"
	"testing"
)

func TestQuerySelect(t *testing.T) {
	db := GetConn()
	defer db.Close()
	rows, err := db.Query("select * from biodata_mahasiswa")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var npm, nama, kelas, jurusan, email string
		err = rows.Scan(&npm, &nama, &kelas, &jurusan, &email)
		if err != nil {
			panic(err)
		}
		fmt.Println("NPM	: ", npm)
		fmt.Println("NAMA 	: ", nama)
		fmt.Println("KELAS 	: ", kelas)
		fmt.Println("JURUSAN	: ", jurusan)
		fmt.Println("EMAIL 	: ", email)
		fmt.Println("")
	}
}
