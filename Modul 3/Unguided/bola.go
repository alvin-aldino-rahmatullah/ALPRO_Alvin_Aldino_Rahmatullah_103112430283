package main

import (
	"fmt"
	"math"
)

func main() {
    var r int
    const pi = 3.1415926535

    // Input jari-jari
    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scan(&r)

    // Menghitung volume bola: (4/3) * pi * r^3
    volume := (4.0 / 3.0) * pi * math.Pow(float64(r), 3)

    // Menghitung luas permukaan bola: 4 * pi * r^2
    luasPermukaan := 4 * pi * math.Pow(float64(r), 2)

    // Menampilkan hasil
    fmt.Printf("Volume bola = %.2f\n", volume)
    fmt.Printf("Luas permukaan bola = %.2f\n", luasPermukaan)
}
