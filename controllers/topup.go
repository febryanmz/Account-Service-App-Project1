package controllers

import (
	"database/sql"
)

func InputNominalTopup(db *sql.DB, idAccount int, nominal int) (int, error) {
	statement, errPrepare := db.Prepare(`INSERT INTO topup (user_id, topup_balance) VALUES (?, ?)`)
	if errPrepare != nil {
		return 0, errPrepare
	}

	_, errSaldo := AddBalance(db, idAccount, nominal)
	if errSaldo != nil {
		return 0, errSaldo
	}

	result, errExec := statement.Exec(&idAccount, &nominal)
	if errExec != nil {
		return 0, errExec
	} else {
		row, _ := result.RowsAffected()
		return int(row), nil
		// if row > 0 {
		// 	fmt.Println("Saldo Bertambah")
		// } else {
		// 	os.Exit(1)
		// 	return int(row), nil // padahal udh return int & error
		// }
	}
	// return 1, nil // hmmmm pusing, minta return lagi

}
