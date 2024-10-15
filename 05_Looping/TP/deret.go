package main

import "fmt"

func main() {
    var n, deret int

    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)

    for angka := 1; angka <= n; angka++ {
        deret += angka
    }

    // Menampilkan hasil
    fmt.Print("Jumlah deret adalah ", deret)
}
