package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetConn() *sql.DB {
	//host=localhost port=5432 dbname=gosql user=loonaticbyu password=Hasooyoung@4 sslmode=disable
	//root:@tcp(localhost:3306)/gosql
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/gosql")
	if err != nil {
		panic(err)
	}
	return db
}

func ErrorHandler(err error) {
	if err != nil {
		fmt.Printf("Terjadi kesalahan: %v\n", err)
	}
}

func InsertQuery(db *sql.DB) {
	var npm, nama, kelas, jurusan, email string
	fmt.Print("Masukkan NPM: ")
	_, err := fmt.Scan(&npm)
	if err != nil {
		ErrorHandler(err)
	}

	fmt.Print("Masukkan Nama: ")
	_, err = fmt.Scan(&nama)
	if err != nil {
		ErrorHandler(err)
	}

	fmt.Print("Masukkan Kelas: ")
	_, err = fmt.Scan(&kelas)
	if err != nil {
		ErrorHandler(err)
	}

	fmt.Print("Masukkan Jurusan: ")
	_, err = fmt.Scan(&jurusan)
	if err != nil {
		ErrorHandler(err)
	}

	fmt.Print("Masukkan Email: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		ErrorHandler(err)
	}

	insertQuery := "INSERT INTO biodata_mahasiswa (npm, nama, kelas, jurusan, email) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(insertQuery, npm, nama, kelas, jurusan, email)
	if err != nil {
		ErrorHandler(err)
	}
	fmt.Println("Data berhasil ditambahkan.")
}

func ReadQuery(db *sql.DB) {
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

func DeleteQuery(db *sql.DB) {
	var npm string
	fmt.Print("Masukkan NPM yang ingin dihapus: ")
	_, err := fmt.Scan(&npm)
	if err != nil {
		ErrorHandler(err)
		return
	}

	deleteQuery := "DELETE FROM biodata_mahasiswa WHERE npm = ?"
	_, err = db.Exec(deleteQuery, npm)
	if err != nil {
		ErrorHandler(err)
		return
	}
	fmt.Println("Data berhasil dihapus.")
}

func UpdateQuery(db *sql.DB) {
	var npm, nama, kelas, jurusan, email string
	fmt.Print("Masukkan NPM yang ingin diperbarui: ")
	_, err := fmt.Scan(&npm)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Print("Masukkan Nama baru: ")
	_, err = fmt.Scan(&nama)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Print("Masukkan Kelas baru: ")
	_, err = fmt.Scan(&kelas)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Print("Masukkan Jurusan baru: ")
	_, err = fmt.Scan(&jurusan)
	if err != nil {
		ErrorHandler(err)
		return
	}

	fmt.Print("Masukkan Email baru: ")
	_, err = fmt.Scan(&email)
	if err != nil {
		ErrorHandler(err)
		return
	}

	updateQuery := "UPDATE biodata_mahasiswa SET nama = ?, kelas = ?, jurusan = ?, email = ? WHERE npm = ?"
	_, err = db.Exec(updateQuery, nama, kelas, jurusan, email, npm)
	if err != nil {
		ErrorHandler(err)
		return
	}
	fmt.Println("Data berhasil diperbarui.")
}

func main() {
	for {
		db := GetConn()
		defer db.Close()

		var choice int
		fmt.Println("Pilih operasi yang ingin dilakukan: ")
		fmt.Println("0. EXIT")
		fmt.Println("1. Insert Data")
		fmt.Println("2. Read Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Print("Masukkan pilihan (0-4): ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			ErrorHandler(err)
			return
		}

		switch choice {
		case 0:
			fmt.Println("Program Selesai.")
			return
		case 1:
			InsertQuery(db)
		case 2:
			ReadQuery(db)
		case 3:
			UpdateQuery(db)
		case 4:
			DeleteQuery(db)
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
