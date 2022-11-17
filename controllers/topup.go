package controllers

import (
	"database/sql"
	"log"
	_entities "project1/entities"
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

// // ------------------------Coba cara INNER JOIN (PUSING JUGA WKWK)------------------------
// func GetTopupHistory(db *sql.DB, idAccount int) ([]_entities.Users, []_entities.TopUp, error) {
// 	query := `
// 	SELECT users.telp, users.firstname, users.lastname, topup.topup_balance, topup.created_at
// 	FROM users
// 	INNER JOIN topup ON topup.user_id
// 	WHERE users.id = ?
// 	`
// 	result, errSelect := db.Query(query, idAccount)
// 	if errSelect != nil { //handling error saat proses menjalankan query
// 		log.Fatal("error select data from query", errSelect.Error())
// 	}
// 	var users []_entities.Users
// 	for result.Next() {
// 		topup := []_entities.TopUp
// 		errScan := result.Scan(
// 			&users.Telp,
// 			&users.Firstname,
// 			&users.Lastname,
// 			&topup.topup_balance,
// 			&topup.created_at,
// 		)
// 		if errScan != nil { // handling ketika ada error pada saat proses scanning
// 			log.Fatal("error scan users & topup", errScan.Error())
// 		}
// 		users.Topup = append(users.Topup, topup)
// 	}

// 	return users, users.Topup, nil
// }
// //-----------------{Pusing WKWKWKWK)---------------------------

func GetHistoryTopUp(db *sql.DB, idAccount int) ([]_entities.TopUp, error) {
	result, errSelect := db.Query("SELECT topup_balance, created_at FROM topup WHERE user_id = ?", &idAccount) // proses menjalankana query SQL
	if errSelect != nil {                                                                                      //handling error saat proses menjalankan query
		log.Fatal("error select ", errSelect.Error())
	}
	var dataUser []_entities.TopUp
	for result.Next() { // membaca tiap baris/row dari hasil query
		var userrow _entities.TopUp                                   // penampung tiap baris data dari db                                                                                                     // membuat variabel penampung
		errScan := result.Scan(&userrow.Balance, &userrow.Created_at) //melakukan scanning data dari masing" row dan menyimpannya kedalam variabel yang dibuat sebelumnya
		if errScan != nil {                                           // handling ketika ada error pada saat proses scannign
			log.Fatal("error scan ", errScan.Error())
		}
		// fmt.Printf("id: %s, nama: %s, email: %s\n", userrow.Id, userrow.Nama, userrow.Email) // menampilkan data hasil pembacaan dari db
		dataUser = append(dataUser, userrow)
	}
	return dataUser, errSelect
}
