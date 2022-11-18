package main

import (
	"fmt"
	"log"
	"os"
	_config "project1/config"
	_controllers "project1/controllers"
)

func main() {
	dbConnection := _config.ConnectToDB()

	defer dbConnection.Close()

	fmt.Printf("---MENU---\n1. Login\n2. Register\n3. Exit\n")
	fmt.Println("Pilih Menu:")
	var menu int
	fmt.Scanln(&menu)
	if menu == 1 {
		//isi syntax login
		//Bryan
		var inTelp string
		var inPass string
		fmt.Println("Masukan No Telp")
		fmt.Scanln(&inTelp)
		fmt.Println("Masukan Password")
		fmt.Scanln(&inPass)
		var idAccount int
		var err error
		idAccount, err = _controllers.GetUserIDbyTelp(dbConnection, inTelp, inPass)
		if err != nil {
			fmt.Println("Login Gagal")
			os.Exit(1) //menghentikan program LANGSUNG (cara baru)
			// log.Fatal("Error gagal login")
		} else {
			fmt.Println("Selamat datang di ALterra Database\nSilahkan dipilih menunya ya kak :)")
		}
		//akhir syntax login1
		var login int
		fmt.Printf("----MENU----\n1. Check Profile\n2. Update Profile\n3. Delete Akun\n4. Cek Profil Orang Lain\n5. TopUp\n6. TopUp History\n7. Transfer\n")
		fmt.Println("Masukkan Pilihan Anda:")
		fmt.Scanln(&login)
		switch login {
		case 1:
			{
				//syntax READ DATA by ID
				//Bryan
				bacaALLData, err := _controllers.GetALLdatabyID(dbConnection, idAccount)
				if err != nil {
					log.Fatal("Error Baca Data")
				}
				for _, v := range bacaALLData {
					fmt.Println("---Data Anda---")
					fmt.Printf("No Telepon: %s\nFirstname: %s\nLastname: %s\nSaldo anda: %d\nDibuat pada: %s\n", v.Telp, v.Firstname, v.Lastname, v.Saldo, v.Created_at.String())
				}
				//----------------
			}
		case 2:
			{
				//syntax Transfer
				//Erlan
				var upid string
				var uptelp string
				var uppass string
				var upfirst string
				var uplast string

				fmt.Println("Masukkan ID user yang akan di update")
				fmt.Scanln(&upid)
				fmt.Println("Masukkan No Telepon user yang akan di update")
				fmt.Scanln(&uptelp)
				fmt.Println("Masukkan Password user yang akan di update")
				fmt.Scanln(&uppass)
				fmt.Println("Masukkan Firstname user yang akan di update")
				fmt.Scanln(&upfirst)
				fmt.Println("Masukkan Lastname user yang akan di update")
				fmt.Scanln(&uplast)

				errPrepare, errExec := _controllers.UpdateUserbyID(dbConnection, uptelp, uppass, upfirst, uplast, upid)
				if errPrepare != nil {
					os.Exit(1)
					if errExec != nil {
						fmt.Println("Update Gagal", err.Error())
						os.Exit(1)
					} else {
						fmt.Println("Update Success")
					}
				}
			}
		case 3:
			{
				//syntax Delete Data by ID
				//Bryan
				err := _controllers.DeleteUserbyID(dbConnection, idAccount)
				if err != nil {
					fmt.Println("Delete Akun Gagal", err.Error())
					os.Exit(1)
				} else {
					fmt.Println("Delete Akun Success")
				}

			}
		case 4:
			{
				//syntax Cek Profil orang lain
				var cekTelp string
				fmt.Println("Masukkan No Telepon user yang akan di cek")
				fmt.Scanln(&cekTelp)
				bacaALLData, err := _controllers.GetALLdatabyTelp(dbConnection, cekTelp)
				if err != nil {
					log.Fatal("Error Baca Data")
				}
				for _, v := range bacaALLData {
					fmt.Println("---Hasil Cek Data---")
					fmt.Printf("Firstname: %s\nLastname: %s\n", v.Firstname, v.Lastname)
				}
			}
		case 5:
			{
				fmt.Println("TopUp")
				//syntax TopUp
				//Bryan
				var nominal int
				fmt.Println("Masukkan Nominal Saldo yang di TopUp :")
				fmt.Scanln(&nominal)
				_, err := _controllers.InputNominalTopup(dbConnection, idAccount, nominal) //input nominal saldo ke tabel top_up
				if err != nil {
					fmt.Println("Top Up Gagal", err.Error())
					os.Exit(1)
				} else {
					fmt.Println("Top Up Berhasil")
				}
			}
		case 6:
			{
				//syntax READ TopUp History by ID
				//Bryan
				//------Coba INNER JOIN----Failed------
				// DataUsers, DataTopUp, err := _controllers.GetTopupHistory(dbConnection, idAccount)
				// if err != nil {
				// 	fmt.Println("Error TopUp History", err.Error())
				// 	os.Exit(1)
				// }
				//___________skip_____________
				//-------Cara Conventional------
				DataUsers, errUsers := _controllers.GetTopUpUsers(dbConnection, idAccount)
				if errUsers != nil {
					log.Fatal("Error Baca Users")
				}
				DataTopUp, errTopUp := _controllers.GetHistoryTopUp(dbConnection, idAccount)
				if errTopUp != nil {
					log.Fatal("Error Baca TopUp")
				}
				for _, u := range DataUsers {
					fmt.Println("---TopUp History---")
					fmt.Printf("No Telepon: %s\nFirstname: %s\nLastname: %s\n", u.Telp, u.Firstname, u.Lastname)
				}
				for _, t := range DataTopUp {
					fmt.Printf("Saldo Topup: Rp.%d\nTopUp pada: %s\n", t.Balance, t.Created_at.String())
				}
			}
		case 7:
			{
				fmt.Println("Transfer")
				//syntax Transfer
				//Erlan
			}
		default:
			{
				fmt.Println("Terima Kasih Sudah Berkunjung di Alterra Immersive Backend 13 :)")
			}

		}

	} else if menu == 2 {
		//syntax Insert into
		//Erlan
		var regtelp string
		var regpass string
		var regfirst string
		var reglast string
		var regsaldo int = 0
		fmt.Println("Masukkan No Telepon anda")
		fmt.Scanln(&regtelp)
		fmt.Println("Masukkan Password anda")
		fmt.Scanln(&regpass)
		fmt.Println("Masukkan Nama Depan anda")
		fmt.Scanln(&regfirst)
		fmt.Println("Masukkan Nama Belakang anda")
		fmt.Scanln(&reglast)

		_, err := _controllers.RegisterUserbyID(dbConnection, regtelp, regpass, regfirst, reglast, regsaldo)
		if err != nil {
			fmt.Println("Register Gagal", err.Error())
			os.Exit(1)
		} else {
			fmt.Println("Register Berhasil")
		}
	} else {
		fmt.Println("Terima Kasih Sudah Berkunjung di Alterra Immersive Backend 13 :)")
	}
}
