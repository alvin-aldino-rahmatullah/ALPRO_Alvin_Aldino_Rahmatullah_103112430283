package main

import "fmt"

func main() {
	var umur int
	var kewarganegaraan string

	fmt.Print("Masukkan umur: ")
	fmt.Scan(&umur)
	fmt.Print("Masukkan kewarganegaraan: ")
	fmt.Scan(&kewarganegaraan)

	if umur >= 17 && kewarganegaraan == "indonesia" {
		fmt.Println("Anda bisa mengikuti pemilu")
	} else if umur < 17 {
		fmt.Println("Anda tidak bisa mengikuti pemilu karena umur Anda belum mencukupi.")
	} else if kewarganegaraan != "indonesia" {
		fmt.Println("Anda tidak bisa mengikuti pemilu karena bukan warga negara indonesia.")
	}
}
