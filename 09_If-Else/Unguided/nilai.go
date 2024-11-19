package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)

	// Jawaban pertanyaan:
	// a. Jika nam = 80.1, output adalah "A" karena nam > 80 terpenuhi.
	// b. Kesalahan program sebelumnya:
	//    1. Tidak menggunakan tanda kutip pada nilai string (seperti A, AB, dll.).
	//    2. Menggunakan variabel 'nam' sebagai string, seharusnya 'nmk'.
	//    3. Tidak menggunakan else-if, menyebabkan beberapa kondisi terpenuhi sekaligus.
	// c. Kode telah diperbaiki dengan pengujian sebagai berikut:
	//    - Input 93.5 → Output: A
	//    - Input 70.6 → Output: B
	//    - Input 49.5 → Output: D
}
