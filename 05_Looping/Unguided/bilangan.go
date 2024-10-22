package main

import "fmt"

func main() {
    var n, hasil int

    fmt.Print("Masukkan bilangan bulat positif n: ")
    fmt.Scan(&n)

    for angka := 1; angka <= n; angka++ {
        hasil += angka
    }

    fmt.Print("Jumlah 1 ditambah sampai ", n,  " adalah ", hasil)
}