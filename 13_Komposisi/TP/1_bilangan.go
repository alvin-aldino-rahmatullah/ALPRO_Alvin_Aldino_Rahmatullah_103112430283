package main

import "fmt"

func main() {
	var n, i, j int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan prima dari 1 hingga", n, "adalah:")

	for i = 2; i <= n; i++ {
		
		for j = 2; j < i; j++ {
			if i%j == 0 {
				break
			}
		}

		if j == i {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
