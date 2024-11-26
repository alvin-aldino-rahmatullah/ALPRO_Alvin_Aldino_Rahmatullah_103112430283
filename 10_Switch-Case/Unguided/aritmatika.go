package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&angka)

	switch {
	case angka%2 != 0:
		fmt.Println("Kategori: Bilangan Ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya", angka, "+", angka+1, "=", angka+(angka+1))

	case angka%2 == 0 && angka%5 != 0 && angka%10 != 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya", angka, "*", angka+1, "=", angka*(angka+1))

	case angka%5 == 0 && angka%10 != 0:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil kuadrat dari", angka, "^2 =", angka*angka)

	case angka%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara", angka, "/ 10 =", float64(angka)/10)

	default:
		fmt.Println("Bilangan tidak sesuai dengan kategori yang ada.")
	}
}
