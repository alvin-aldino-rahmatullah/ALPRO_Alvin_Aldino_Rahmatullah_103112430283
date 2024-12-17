package main

import "fmt"

func main() {
	var bil int

	fmt.Print("Masukkan angka : ")
	fmt.Scan(&bil)

	for j := 1; j <= bil; j++ {
		if bil%j == 0 {
			fmt.Print(j, "")
		}
	}
	fmt.Println()
}
