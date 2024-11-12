package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&huruf)

	if huruf == "A" ||huruf == "a" || huruf == "I" || huruf == "i" || huruf == "U" || huruf == "u"  || huruf == "E" || huruf == "e" || huruf == "O" || huruf == "o" {
		fmt.Println("Huruf Vokal")
	} else {
		fmt.Println("Huruf Konsonan")
	}
}
