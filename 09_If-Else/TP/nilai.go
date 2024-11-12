package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan Nilai : ")
	fmt.Scan(&nilai)

	if nilai > 90 {
		fmt.Println("Selamat nilai kamu A")
	} else if nilai >= 80 && nilai <= 90 {
		fmt.Println("Selamat nilai kamu AB")
	} else if nilai >= 70 && nilai <= 80 {
		fmt.Println("Selamat nilai kamu B")
	} else if nilai < 70 {
		fmt.Println("Belajar lagi nilai kamu C")
	}
}

