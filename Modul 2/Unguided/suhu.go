package main

import "fmt"

func main() {
	// Membuat variabel fahrenheit dan celsius
	var fahrenheit, celsius float32

	// Meminta input suhu dalam Fahrenheit dari user
	fmt.Print("Masukkan suhu dalam derajat Fahrenheit: ")
    // Scan menyimpan data input user yang telah di isi ke dalam variabel fahrenheit
	fmt.Scanln(&fahrenheit)

	// Melakukan konversi dari Fahrenheit ke celcius:
	celsius = (fahrenheit - 32) * 5 / 9

	// Menampilkan hasil konversi fahrenheit ke celcius
	fmt.Println("Suhu dalam derajat celsius = ", celsius)
}