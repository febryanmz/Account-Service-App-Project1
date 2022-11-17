package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

func GetUserIDbyTelp(db *sql.DB, inTelp string, inPass string) (int, error) { // value disini
	results := db.QueryRow("SELECT id from users where telp = ? && pass = ?", &inTelp, &inPass) // ngebaca data input telp & pass
	var idAccount int
	err := results.Scan(&idAccount)

	if err != nil {
		return 0, err //package errors (searching)
	}
	return idAccount, err
}

func RegisterUserbyID(db *sql.DB, regtelp string, regpass string, regfir string, reqpas string) (error, error) {
	statement, errPrepare := db.Prepare("INSERT INTO users (telp, pass, firstname, lastname) VALUES (?, ?, ?, ? )")
	if errPrepare != nil {
		return nil, errPrepare
	}
	result, errExec := statement.Exec(&regtelp, &regpass, &regfir, &reqpas)
	if errExec != nil {
		log.Fatal("error exec register", errExec.Error())

	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("Register Berhasil")

		} else {
			os.Exit(1)
		}

	}
	return errPrepare, errExec
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
