package main

import (
	"fmt"
	"log"

	_config "project1/config"
	_entities "project1/entities"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbConnection := _config.ConnectToDB()

	defer dbConnection.Close() // menutup koneksi

	//buat mekanisme menu
	fmt.Printf("MENU:\n1. Login\n2. Register\n3. Cek Profil\n4. Update Profil\n5. Delete Profil\n6. TopUp\n7. Transfer\n")
	fmt.Println("Masukkan pilihan anda:")
	var pilihan int
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		{

		}
	case 2:
		{
			newUser := _entities.Users{}
			fmt.Println("Masukkan No Telepon anda")
			fmt.Scanln(&newUser.Telp)
			fmt.Println("Masukkan Password anda")
			fmt.Scanln(&newUser.Pass)
			fmt.Println("Masukkan Nama Depan anda")
			fmt.Scanln(&newUser.Firstname)
			fmt.Println("Masukkan Nama Belakang anda")
			fmt.Scanln(&newUser.Lastname)
			newUser.Saldo = 0

			statement, errPrepare := dbConnection.Prepare(`INSERT INTO users (telp, pass, firstname, lastname, saldo) VALUES (?, ?, ?, ?, ?)`)
			if errPrepare != nil {
				log.Fatal("error prepare register", errPrepare.Error())
			}

			result, errExec := statement.Exec(newUser.Telp, newUser.Pass, newUser.Firstname, newUser.Lastname, newUser.Saldo)
			if errExec != nil {
				log.Fatal("error exec register", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("Insert register")
				} else {
					fmt.Println("Insert register")
				}
			}
		}
	case 3:
		{
			// fmt.Println("baca data by ID")
			bacaUser := _entities.Users{}
			fmt.Println("Masukkan ID User")
			fmt.Scanln(&bacaUser.Id)

			results := dbConnection.QueryRow("SELECT telp, firstname, lastname, saldo, created_at from users where id = ?", &bacaUser.Id)
			var dataUser _entities.Users
			err := results.Scan(&dataUser.Telp, &dataUser.Firstname, &dataUser.Lastname, &dataUser.Saldo, &dataUser.Created_at)

			if err != nil {
				log.Fatal("error select ", err.Error())
			}
			fmt.Printf("id: %d\ntelp: %s\nfirstname: %s\nlastname: %s\nsaldo: %d\ncreated_at: %s\n", dataUser.Id, dataUser.Telp, dataUser.Firstname, dataUser.Lastname, dataUser.Saldo, dataUser.Created_at.String())
			//----------------

		}
	case 4:
		{
			// (update user by ID)
			updateUser := _entities.Users{}
			fmt.Println("Masukkan ID user yang akan di update")
			fmt.Scanln(&updateUser.Id)
			fmt.Println("Masukkan No Telepon user yang akan di update")
			fmt.Scanln(&updateUser.Telp)
			fmt.Println("Masukkan Password user yang akan di update")
			fmt.Scanln(&updateUser.Pass)
			fmt.Println("Masukkan Firstname user yang akan di update")
			fmt.Scanln(&updateUser.Firstname)
			fmt.Println("Masukkan Lastname user yang akan di update")
			fmt.Scanln(&updateUser.Lastname)

			statement, errPrepare := dbConnection.Prepare(`UPDATE users set telp = ?, pass = ?, firstname = ?, lastname = ? where Id = ?`)
			if errPrepare != nil {
				log.Fatal("error prepare update", errPrepare.Error())
			}

			result, errExec := statement.Exec(updateUser.Pass, updateUser.Firstname, updateUser.Lastname, updateUser.Telp, updateUser.Id)
			if errExec != nil {
				log.Fatal("error exec update", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("update berhasil")
				} else {
					fmt.Println("update gagal")
				}
			}

		}
	case 5:
		{
			// fmt.Println("delete")
			deleteUser := _entities.Users{}
			fmt.Println("Masukkan id user yang akan di DELETE")
			fmt.Scanln(&deleteUser.Id)

			statement, errPrepare := dbConnection.Prepare(`DELETE from users where id = ?`)
			if errPrepare != nil {
				log.Fatal("error prepare delete", errPrepare.Error())
			}

			result, errExec := statement.Exec(deleteUser.Id)
			if errExec != nil {
				log.Fatal("error exec delete", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("delete berhasil")
				} else {
					fmt.Println("delete gagal")
				}
			}
		}
	case 6:
		{
			// var id_pilihan int
			// -----Topup Saldo-----
			Topup := _entities.TopUp{}
			fmt.Println("Masukkan Id yang ingin di TopUp")
			fmt.Scanln(&Topup.Id)

			fmt.Println("Masukkan Nominal TopUp")
			fmt.Scanln(&Topup.Balance)

			statement, errPrepare := dbConnection.Prepare(`INSERT INTO topup (id, balance) VALUES (?, ?)`)
			if errPrepare != nil {
				log.Fatal("error prepare topup", errPrepare.Error())
			}

			result, errExec := statement.Exec(Topup.Balance)
			if errExec != nil {
				log.Fatal("error exec topup", errExec.Error())
			} else {
				row, _ := result.RowsAffected()
				if row > 0 {
					fmt.Println("Topup berhasil")
				} else {
					fmt.Println("Topup gagal")
				}
			}
			// -----Baca Saldo by Id------

			// -----Update Saldo di Users-----

			// topupSaldo =
			// statement, errPrepare := dbConnection.Prepare(`UPDATE users set saldo = ? where id = ?`)
			// if errPrepare != nil {
			// 	log.Fatal("error prepare update", errPrepare.Error())
			// }

			// result, errExec := statement.Exec(topupSaldo.Saldo, topupSaldo.Id)
			// if errExec != nil {
			// 	log.Fatal("error exec update", errExec.Error())
			// } else {
			// 	row, _ := result.RowsAffected()
			// 	if row > 0 {
			// 		fmt.Println("update berhasil")
			// 	} else {
			// 		fmt.Println("update gagal")
			// 	}
			// }

		}
	case 7:
		{
			fmt.Print("Transfer")
		}

	}
}
