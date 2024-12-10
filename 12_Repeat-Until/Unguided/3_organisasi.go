package main

import "fmt"

func main() {
	var target, donasi, total, donatur int

	fmt.Scan(&target)

	for done := false; !done; {

		fmt.Scan(&donasi)
		total += donasi
		donatur++

		fmt.Println("Donatur", donatur, ": Menyumbang ", donasi, " Total terkumpul ", total)
		done = total >= target
	}

	fmt.Println("Target tercapai! Total donasi: ", total, " dari ", donatur, " donatur ")
}
