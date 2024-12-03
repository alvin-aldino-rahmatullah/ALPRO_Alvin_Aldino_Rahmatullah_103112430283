package main

import "fmt"

func main() {
	var guess int


    for {
        fmt.Print("Tebak angka (1-10): ")
        fmt.Scan(&guess)

        if guess == 7 { 
            fmt.Println("Selamat, tebakan Anda benar!")
            break
        }
        fmt.Println("Tebakan Anda salah, coba lagi.") 
    }
}
