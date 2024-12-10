package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print(" Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	digit := 0
	for done := false; !done; {
		bilangan /= 10
		digit++
		done = bilangan == 0
	}

	fmt.Print(" Banyaknya digit adalah: ", digit)
}
