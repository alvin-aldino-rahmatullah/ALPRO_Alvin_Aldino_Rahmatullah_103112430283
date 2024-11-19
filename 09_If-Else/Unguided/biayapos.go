package main

import "fmt"

func main() {
	var berat int
	fmt.Print("Masukkan berat parsel (gram): ")
	fmt.Scanln(&berat)


	beratKg := berat / 1000
	sisaGram := berat % 1000
	biayaDasar := beratKg * 10000

	biayaTambahan := 0
	if sisaGram > 0 {
		if sisaGram < 500 {
			biayaTambahan = sisaGram * 5
		} else {
			biayaTambahan = sisaGram * 15
		}
	}

	if beratKg < 10 && sisaGram >= 1000 {
		biayaTambahan = 0
	}

	
	totalBiaya := biayaDasar + biayaTambahan

	fmt.Print("Berat parsel:", berat, "gram")
	fmt.Print("Detail berat: ", beratKg, "kg", sisaGram, "gram")
	fmt.Print("Detail biaya: Rp.", biayaDasar, "+ Rp.",biayaTambahan)
	fmt.Print("Total biaya: Rp", totalBiaya)
}
