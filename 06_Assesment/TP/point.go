package main

import (
	"fmt"
)

func main() {
	var jumlahbarang int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahbarang)

	totalpoint := 0

	for i := 1; i <= jumlahbarang; i++ {
		totalpoint += 10
		if i > 5 {
			totalpoint += 5
		}
	}
	fmt.Print("Total point yang di dapat ", totalpoint, " point ")
}