package main

import "fmt"

func main() {
    //Membuat variabel sisi dan volume
    var sisi , volume float64
    //Meminta input panjang sisi kubus yang akan di isi oleh user
    fmt.Print(" Masukkan Panjang Sisi Kubus = " )
	//Scan membaca sisi yang telah di isi oleh user lalu menyimpannya ke variabel sisi
    fmt.Scan(&sisi)
    //Menghitung variabel volume dengan cara menggabunggkan variabel sisi dengan rumus volume kubus
    volume = sisi * sisi * sisi
	//Menampilkan volume yang telah di hitung
    fmt.Println(" Volume Kubus adalah = ", volume)
}
