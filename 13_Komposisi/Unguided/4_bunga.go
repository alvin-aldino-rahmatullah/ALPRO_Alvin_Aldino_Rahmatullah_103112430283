package main

import "fmt"


func main() {
	var bunga, pita string
	var jumlahBunga int

	for {
		fmt.Print("Bunga: ")
		fmt.Scan(&bunga)

		if bunga == "selesai" {
			break
		}

		if jumlahBunga > 0 {
			pita += " – " + bunga
		} else {
			pita += bunga
		}

		jumlahBunga++
	}

	
	fmt.Println("Pita:", pita)
	fmt.Println("Jumlah bunga:", jumlahBunga)
}
