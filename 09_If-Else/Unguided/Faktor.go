package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	
	fmt.Print("Faktor: ")
	prima := b > 1
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i)
			if i != 1 && i != b {
				prima = false
			}
		}
	}
	fmt.Println()

	
	fmt.Print( prima)
}
