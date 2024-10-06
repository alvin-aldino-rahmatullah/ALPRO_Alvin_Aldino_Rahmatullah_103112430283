package main

import (
	"fmt"
	"math"
)

func main() {
	// Membuat variabel jarijari
    var jarijari float64

    // Meminta user untuk memasukkan jari jari liangkaran 
    fmt.Print("Masukkan jari-jari lingkaran: ")
	// Scan menyimpan data yang telah di isi user ke dalam variabel jarijari
    fmt.Scanln(&jarijari)

    // Menghitung luas lingkaran
    luas := math.Pi * jarijari * jarijari

    // Menampilkan hasil perhitungan luas lingkaran
    fmt.Printf("Luas lingkaran dengan jari-jari %.1f adalah %.1f", jarijari, luas)
}
