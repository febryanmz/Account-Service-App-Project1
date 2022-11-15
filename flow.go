package main

import "fmt"

func main() {

	fmt.Printf("MENU:\n1. Login\n2. Register\n3. Exit\n")
	fmt.Println("Selamat Datang di Alterra Immersive Backend 13 :)\nPilih Menu:")
	var menu int
	fmt.Scanln(&menu)
	if menu == 1 {
		//isi syntax login

		//akhir syntax login
		fmt.Println("Proceed to login")
		var login int
		fmt.Printf("MENU:\n1. Cek Profil\n2. Update Profil\n3. Delete Profil\n4. TopUp\n5. Transfer\n")
		fmt.Println("Masukkan Pilihan Anda:")
		fmt.Scanln(&login)
		switch login {
		case 1:
			{
				fmt.Println("Cek Profil")
				//syntax READ DATA by ID
			}
		case 2:
			{
				fmt.Println("Update Profil")
				//syntax Update Data by ID
			}
		case 3:
			{
				fmt.Println("Delete Profil")
				//syntax Delete Data by ID
			}
		case 4:
			{
				fmt.Println("Top Up")
				//syntax TopUp
			}
		case 5:
			{
				fmt.Println("Transfer")
				//syntax Transfer
			}

		}

	} else if menu == 2 {
		fmt.Println("Proceed to register")
	} else {
		fmt.Println("Terima Kasih Sudah Berkunjung di Alterra Immersive Backend 13 :)")
	}
}
