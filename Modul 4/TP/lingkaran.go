package main

import (
	"fmt"
	"math"
)

func main() {
    // Membuat variabel jarijari 
    var jariJari float64
    // Meminta user memasukkan jari jari lingkaran
       fmt.Print("Masukkan jari-jari lingkaran: ")
    // Scan menyimpan variabel jari jari yang telah di isi user
    fmt.Scanln(&jariJari)
    
    // Menghitung luas dengan rumus
    // math.Pi yang menyimpan nilai Pi ≈ 3.14
    luas := math.Pi * jariJari * jariJari
    // Menghitung luas dengan rumus dan memasukkan variabel jarijari
    keliling := 2 * math.Pi * jariJari

    // Menampilkan hasil yang telah di hitung
    // Menambahkan code %.2f untuk menampillkan 2 angka dibelakang koma
    fmt.Printf("Luas lingkaran: %.2f\n", luas)
    fmt.Printf("Keliling lingkaran: %.2f\n", keliling)
}
