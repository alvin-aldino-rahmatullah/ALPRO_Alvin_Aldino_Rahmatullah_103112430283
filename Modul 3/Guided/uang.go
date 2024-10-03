package main

import "fmt"

func main () {
	//Membuat variabel rupiah dan hasil
	var rupiah, hasil float32
    //Meminta input rupiah yang akan di isi oleh user
	fmt.Print("Masukkan Nilai Rupiah = ")
	//Scan membaca rupiah yang telah di isi oleh user lalu menyimpannya ke variabel rupiah
	fmt.Scan (&rupiah)
    //menghitung variabel rupiah dengan membagi dengan 15000
	hasil = rupiah / 15000
	//Menampilkan hasil yang telah di hitung
	fmt.Print("konversi Dolar = ", hasil)
}