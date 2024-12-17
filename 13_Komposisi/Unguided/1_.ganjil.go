package main

import "fmt"

func main() {
	var n, ganjil int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	ganjil = 0

	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			ganjil++
		}
	}
	fmt.Println("Terdapat ", ganjil, " bilangan ganjil")
}
