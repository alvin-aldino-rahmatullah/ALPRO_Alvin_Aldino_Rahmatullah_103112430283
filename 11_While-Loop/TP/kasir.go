package main

import "fmt"

func main() {
	var totalBelanja, hargaBarang float64  
	var namaBarang string       

	for {
		fmt.Print("Masukkan nama barang atau ketik 'selesai' untuk mengakhiri: ")
		fmt.Scanln(&namaBarang) 

		if namaBarang == "selesai" {
			break 
		}

		fmt.Print("Masukkan harga barang: ")
		fmt.Scanln(&hargaBarang) 

		totalBelanja += hargaBarang 
		fmt.Print("Barang ", namaBarang, " dengan harga ", hargaBarang, " telah ditambahkan ke total belanja. ")
	}
	fmt.Print("Total belanja Anda adalah:", totalBelanja)
}
