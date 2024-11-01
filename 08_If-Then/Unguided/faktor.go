package main

import "fmt"

func main() {
    var x, y int
	var FaktorX, FaktorY bool

    fmt.Print("Masukkan dua bilangan bulat positif (x y): ")
    fmt.Scan(&x, &y)

    FaktorY = false
    FaktorX = false

    if y % x == 0 {
        FaktorY = true
    }

    
    if x % y == 0 {
        FaktorX = true
    }

    fmt.Println(FaktorY)
    fmt.Println(FaktorX)
}
