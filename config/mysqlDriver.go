package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	var connectionString = "root:nephilim17@tcp(127.0.0.1:3306)/alterra_project1?parseTime=true"
	db, err := sql.Open("mysql", connectionString) // membuka koneksi ke daatabase
	if err != nil {                                // pengecekan error yang terjadi ketika proses open connection
		log.Fatal("error open connection", err.Error())
	}

	errPing := db.Ping() // mengecek apakah apliasi masih terkoneksi ke database
	if errPing != nil {  //handling error ketika gagal konek ke db
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("koneksi berhasil")
	}

	return db
}
