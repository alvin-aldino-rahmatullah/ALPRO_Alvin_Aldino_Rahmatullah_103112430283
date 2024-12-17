package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan, totalBerat float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		if beratKiri < 0 || beratKanan < 0 {
			break
		}


		totalBerat = beratKiri + beratKanan

		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(beratKiri - beratKanan)

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}
}
