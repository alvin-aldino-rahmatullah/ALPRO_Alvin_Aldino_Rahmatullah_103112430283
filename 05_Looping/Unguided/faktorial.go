package main

import (
    "fmt"
)

func main() {
    var bilangan int
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)

    hasilFaktorial := 1
    for i := 1; i <= bilangan; i++ {
        hasilFaktorial *= i
    }
    fmt.Println("Hasil faktorial dari", bilangan, "adalah:", hasilFaktorial)
}
