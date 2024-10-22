package main

import (
    "fmt"
)

func main() {
    var bilangan, pangkat int
    fmt.Print("Masukkan bilangan dasar : ")
    fmt.Scan(&bilangan)
	
    fmt.Print("Masukkan pangkat : ")
    fmt.Scan(&pangkat)

    hasilPangkat := 1

    for i := 0; i < pangkat; i++ {
        hasilPangkat *= bilangan
    }

    fmt.Println("Hasil pemangkatan adalah:", hasilPangkat)
}
