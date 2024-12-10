package main

import "fmt"

func main() {
    var n, total int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&n)

    for i := 1; i < n; i++ {
        if n%i == 0 { 
            total += i
        }
    }

    if total == n {
        fmt.Println("Ya") 
    } else {
        fmt.Println("Tidak") 
    }
}
