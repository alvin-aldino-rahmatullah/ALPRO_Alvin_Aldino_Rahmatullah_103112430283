package main

import (
	"fmt"
)

func main () {
	var n int
	fmt.Print("Masukkan bilangan bulat : ")
	fmt.Scan(&n)

	fmt.Println(" Hasil kuadrat dari 1 sampai", n)

	for i := 1; i <= n; i++ {
		hasilKuadrat := i * i
		fmt.Printf("%d ", hasilKuadrat )
	}
}