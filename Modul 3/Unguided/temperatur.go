package main

import "fmt"

func main() {
    var celsius float64

    // Membaca input dari pengguna
    fmt.Print("Masukkan temperatur dalam Celsius: ")
    fmt.Scan(&celsius)

    // Menghitung Fahrenheit
    fahrenheit := (celsius * 9 / 5) + 32

    // Menghitung Reamur
    reamur := celsius * 4 / 5

    // Menghitung Kelvin
    kelvin := (fahrenheit + 459.67) * 5 / 9

    // Menampilkan hasil
    fmt.Printf("Temperatur dalam Fahrenheit: %.2f°F\n", fahrenheit)
    fmt.Printf("Temperatur dalam Reamur: %.2f°R\n", reamur)
    fmt.Printf("Temperatur dalam Kelvin: %.2f K\n", kelvin)
}
