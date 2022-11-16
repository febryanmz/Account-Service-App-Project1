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
		fmt.Printf("----MENU----\n1. Check Profile\n2. Update Profile\n3. Delete Akun\n4. TopUp\n5. Transfer\n")
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
				fmt.Println("Update Profile")
				//syntax Update Data by ID
				//Erlan
			}
		case 3:
			{
				//syntax Delete Data by ID
				//Bryan
				err := _controllers.DeleteUserbyID(dbConnection, idAccount)
				if err != nil {
					os.Exit(1)
				} else {
					fmt.Println("Success")
				}

			}
		case 4:
			{
				fmt.Println("Top Up")
				//syntax TopUp
				//Bryan
			}
		case 5:
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
		fmt.Println("Proceed to register")
		//syntax Insert into
		//Erlan
	} else {
		fmt.Println("Terima Kasih Sudah Berkunjung di Alterra Immersive Backend 13 :)")
	}
}
