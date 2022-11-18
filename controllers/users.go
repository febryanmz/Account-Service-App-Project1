package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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

func DeleteUserbyID(db *sql.DB, idAccount int) error {
	statement, errPrepare := db.Prepare(`DELETE from users where id = ?`)
	if errPrepare != nil {
		log.Fatal("Error prepare Delete", errPrepare.Error())
	}

	result, errExec := statement.Exec(&idAccount)
	if errExec != nil {
		log.Fatal("Error exec Delete", errExec.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Delete Berhasil")
		} else {
			os.Exit(1)
		}
	}
	return errExec
}
func GetSaldo(db *sql.DB, idAccount int) (int, error) {
	results := db.QueryRow("SELECT saldo FROM users WHERE id = ?", &idAccount) // ngebaca data input telp & pass
	var saldo int
	err := results.Scan(&saldo)
	if err != nil {
		return 0, err //package errors (searching)
	}
	return saldo, err
}

func AddBalance(db *sql.DB, idAccount int, nominal int) (int, error) {
	var query = "update users set saldo = (?) where id = (?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		return 0, errPrepare
	}
	saldo, errSaldo := GetSaldo(db, idAccount)
	if errSaldo != nil {
		return 0, errSaldo
	}
	var saldoTopUp = nominal + saldo
	result, err := statement.Exec(&saldoTopUp, &idAccount)
	if err != nil {
		return 0, err
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
		//--------------
		// if err != nil {
		//     return -1, err
		// } else {
		//     return saldoTopUp, nil
		// }
		//---------------
		// if row == 0 {
		//     os.Exit(1)
		// } else {
		//     return int(row), nil
		// }
	}
	// return saldoTopUp, nil
}
func RegisterUserbyID(db *sql.DB, regtelp string, regpass string, regfirst string, reglast string, regsaldo int) (int, error) {
	statement, errPrepare := db.Prepare("INSERT INTO users (telp, pass, firstname, lastname, saldo) VALUES (?, ?, ?, ?, ? )")
	if errPrepare != nil {
		return 0, errPrepare
	}
	result, errExec := statement.Exec(&regtelp, &regpass, &regfirst, &reglast, &regsaldo)
	if errExec != nil {
		return 0, errExec

	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
	}

}

func UpdateUserbyID(db *sql.DB, uptelp string, uppass string, upfirst string, uplast string, upid string) (error, error) {
	statement, errPrepare := db.Prepare(`UPDATE users set telp = ?, pass = ?, firstname = ?, lastname = ? where Id = ?`)
	if errPrepare != nil {
		return nil, errPrepare
	}
	result, errExec := statement.Exec(&uptelp, &uppass, &upfirst, &uplast, &upid)
	if errExec != nil {
		log.Fatal("error exec register", errExec.Error())

	} else {
		row, _ := result.RowsAffected()

		if row > 0 {
			fmt.Println("Update Berhasil")

		} else {
			os.Exit(1)
		}

	}
	return errPrepare, errExec
}

func GetTopUpUsers(db *sql.DB, idAccount int) ([]_entities.Users, error) {
	result, errSelect := db.Query("SELECT telp, firstname, lastname from users where id = ?", &idAccount) // proses menjalankana query SQL
	if errSelect != nil {                                                                                 //handling error saat proses menjalankan query
		log.Fatal("error select ", errSelect.Error())
	}
	var dataUser []_entities.Users
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow _entities.Users                                                  // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
		errScan := result.Scan(&userrow.Telp, &userrow.Firstname, &userrow.Lastname) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                                          // handling ketika ada error pada saat proses scannign
			log.Fatal("error scan ", errScan.Error())
		}
		// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
		dataUser = append(dataUser, userrow)
	}
	return dataUser, errSelect
}

func GetALLdatabyTelp(db *sql.DB, cekTelp string) ([]_entities.Users, error) {
	result, errSelect := db.Query("SELECT firstname, lastname from users where telp = ?", &cekTelp) // proses menjalankana query SQL
	if errSelect != nil {                                                                           //handling error saat proses menjalankan query
		log.Fatal("error select ", errSelect.Error())
	}
	var dataUser []_entities.Users
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow _entities.Users                                   // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
		errScan := result.Scan(&userrow.Firstname, &userrow.Lastname) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                           // handling ketika ada error pada saat proses scannign
			log.Fatal("error scan ", errScan.Error())
		}
		// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
		dataUser = append(dataUser, userrow)
	}
	return dataUser, errSelect
}
