package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1 
    var guess int
    maxAttempts := 5

    fmt.Println("Selamat datang di permainan tebak angka!")
    fmt.Println("Saya telah memilih angka antara 1 hingga 100.")
    fmt.Println("Anda memiliki", maxAttempts, "kesempatan untuk menebak.")

    for attempts := 1; attempts <= maxAttempts; attempts++ {
        fmt.Printf("Tebakan ke-%d: ", attempts)
        fmt.Scan(&guess)

        if guess < target {
            fmt.Println("Tebakan Anda terlalu kecil.")
        } else if guess > target {
            fmt.Println("Tebakan Anda terlalu besar.")
        } else {
            fmt.Println("Selamat! Anda telah menebak angka yang benar:", target)
            return 
        }
    }

    fmt.Println("Maaf, Anda telah menggunakan semua kesempatan. Angka yang benar adalah:", target)
}
