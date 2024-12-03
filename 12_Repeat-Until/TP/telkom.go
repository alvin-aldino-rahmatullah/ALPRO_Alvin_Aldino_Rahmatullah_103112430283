package main

import "fmt"

func main() {
    var input string

    for {
        fmt.Print("Masukkan kata: ")
        fmt.Scanln(&input)

        if input == "telkom" { 
            fmt.Println("Program selesai")
            break
        }
        fmt.Println("Anda mengetik:", input) 
    }
}
