package main

import "fmt"

func main () {
	//Membuat variabel nilai dan hasil
	var nilai, hasil int
    //Meminta input nilai yang akan di isi oleh user
	fmt.Print(" Masukkan nilai " )
	//Scan membaca nilai yang telah di isi oleh user lalu menyimpannya ke variabel nilai
	fmt.Scan (&nilai)
    //Menghitung variabel hasil dengan cara menggabunggkan variabel nilai dengan rumus yang ada
	hasil = (2/(nilai+5)+5)
	//Menampilkan hasil yang telah di hitung
	fmt.Print("hasil = ", hasil)
}