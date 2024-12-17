package main

import "fmt"

func main() {
	var n int
	var prima bool

	fmt.Scan(&n)

	prima = true

	for i := 2; i < n; i++ {
		if n%i == 0 {
			prima = false
			break
		}
	}

	if prima {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
