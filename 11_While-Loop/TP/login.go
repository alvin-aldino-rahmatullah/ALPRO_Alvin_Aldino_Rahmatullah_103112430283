package main

import "fmt"

func main() {
	const password = "12345" 
	var input string         
	       

	for i := 0; i < 3; i++ {
		fmt.Print("Masukkan password: ")
		fmt.Scan(&input) 

		if input == password {
			fmt.Println("Login berhasil!")
			return
		} else {
			fmt.Println("Password salah.")
		}
	}

	fmt.Println("Login ditolak")
}
