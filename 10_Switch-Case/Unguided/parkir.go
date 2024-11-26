package main

import "fmt"

func main() {
    var jenisKendaraan string
    var durasi, tarifPerJam, totalBiaya int

    fmt.Print("Masukkan jenis kendaraan (motor, mobil, truk): ")
    fmt.Scan(&jenisKendaraan)


    fmt.Print("Masukkan durasi parkir dalam jam: ")
    fmt.Scan(&durasi)


    switch jenisKendaraan {
    case "motor":
        tarifPerJam = 2000
    case "mobil":
        tarifPerJam = 5000
    case "truk":
        tarifPerJam = 8000
    default:
        fmt.Println("Jenis kendaraan tidak valid")
    }

    totalBiaya = tarifPerJam * durasi

    fmt.Print("Total Biaya Parkir: Rp ", totalBiaya)
}
