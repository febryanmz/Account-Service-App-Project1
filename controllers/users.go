package controllers

import (
	"database/sql"
	"log"
	_entities "project1/entities"
)

func GetALLdatabyID(db *sql.DB, idAccount int) ([]_entities.Users, error) {
	result, errSelect := db.Query("SELECT telp, firstname, lastname, saldo, created_at from users where id = ?", &idAccount) // proses menjalankana query SQL
	if errSelect != nil {                                                                                                    //handling error saat proses menjalankan query
		log.Fatal("error select ", errSelect.Error())
	}
	var dataUser []_entities.Users
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow _entities.Users                                                                                       // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
		errScan := result.Scan(&userrow.Telp, &userrow.Firstname, &userrow.Lastname, &userrow.Saldo, &userrow.Created_at) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                                                                               // handling ketika ada error pada saat proses scannign
			log.Fatal("error scan ", errScan.Error())
		}
		// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
		dataUser = append(dataUser, userrow)
	}
	return dataUser, errSelect

}

func GetUserIDbyTelp(db *sql.DB, inTelp string, inPass string) (int, error) { // value disini
	results := db.QueryRow("SELECT id from users where telp = ? && pass = ?", &inTelp, &inPass) // ngebaca data input telp & pass
	var idAccount int
	err := results.Scan(&idAccount)

	if err != nil {
		return 0, err //package errors (searching)
	}
	return idAccount, err
}
