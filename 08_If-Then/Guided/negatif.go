package main

import "fmt"

func main() {
    var nilai int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&nilai)

    if nilai < 0 {
        fmt.Print("true")
    } 
	fmt.Print("false")
	
    
}