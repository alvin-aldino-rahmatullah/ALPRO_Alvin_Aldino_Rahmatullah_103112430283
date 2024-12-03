package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&angka)

	for angka > 0 {
		digit := angka % 10
		fmt.Println(digit)
		angka /= 10
	}

}
