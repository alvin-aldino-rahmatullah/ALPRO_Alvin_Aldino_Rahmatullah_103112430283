package main

import "fmt"

func main() {
	//Membuat variabel alas, tinggi dan hasil
    var alas, tinggi, hasil float32
     
	//Meminta input alas yang akan di isi oleh user
    fmt.Print(" Masukkan nilai alas = ") 
	//Scan membaca alas yang telah di isi oleh user lalu menyimpannya ke variabel alas
    fmt.Scan(&alas)
	//Meminta input tinggi yang akan di isi oleh user
    fmt.Print(" Masukkan nilai tinggi = ")
	//Scan membaca tinggi yang telah di isi oleh user lalu menyimpannya ke variabel alas
    fmt.Scan(&tinggi)

	//Menghitung variabel hasil dengan cara menggabunggkan variabel alas dan tinggi dengan rumus luas segitga
    hasil = 0.5 * alas * tinggi
	//Menampilkan hasil yang telah di hitung
    fmt.Println(" Luas Segitiga = ", hasil)
}
