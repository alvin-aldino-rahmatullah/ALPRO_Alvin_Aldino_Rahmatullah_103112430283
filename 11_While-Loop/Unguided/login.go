package main

import "fmt"

func main() {
	var username, password string

	const UsernameAsli = "Admin"
	const PasswordAsli = "Admin"
	attempts := 0

	for {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)

		if username == UsernameAsli && password == PasswordAsli {
			break
		} else {
			fmt.Println("Username atau password salah. Silakan coba lagi.")
			attempts++
		}
	}

	fmt.Print("Login berhasil setelah ", attempts, "  percobaan gagal.")
}
