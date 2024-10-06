package main

import "fmt"

func main() {
	//memasukkan variabel bernama tahun
    var tahun int

    // Input tahun dari pengguna
    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&tahun)

    // Memeriksa apakah tahun kabisat
    if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
        fmt.Println("Tahun", tahun, "adalah tahun kabisat")
    } else {
        fmt.Println("Tahun", tahun, "bukanlah tahun kabisat") 
    }
}
