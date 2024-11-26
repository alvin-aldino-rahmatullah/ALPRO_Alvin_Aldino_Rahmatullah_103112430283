package main
import "fmt"

func main() {
	
	var waktu int
	fmt.Print("Masukkan waktu: ")
	fmt.Scan(&waktu)
	
	switch {
	case waktu >= 0 && waktu < 12:
		fmt.Println(waktu, "am")
	case waktu >= 12 && waktu <= 24:
		waktu -= 12
		fmt.Println(waktu, "pm")	
	default : 
	 fmt.Print("Waktu yang anda masukkan tidak tersedia")
	 
	}
}
