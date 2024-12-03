package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan bilangan x : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan y : ")
	fmt.Scan(&y)

	hasil := 0
	for x >= y {
		x -= y
		hasil++
	}

	fmt.Print(hasil)
}
