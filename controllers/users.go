package controllers

import (
	"database/sql"
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
