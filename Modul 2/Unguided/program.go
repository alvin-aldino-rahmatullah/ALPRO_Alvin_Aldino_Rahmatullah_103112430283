package main

import "fmt"
func main() {
 var (
// Membuat variabel satu, dua dan tiga
 satu, dua, tiga string
 temp string
 )
 // Meminta user untuk memasukkan input pertama
 fmt.Print("Masukan input string: ")
 // Scan menyimpan data input pertama yang telah di isi user ke dalam variabel satu
 fmt.Scanln(&satu)
 
 // Meminta user untuk memasukkan input kedua
 fmt.Print("Masukan input string: ")
 // Scan menyimpan data input kedua yang telah di isi user ke dalam variabel dua
 fmt.Scanln(&dua)
 
 // Meminta user untuk memasukkan input ketiga
 fmt.Print("Masukan input string: ")
 // Scan menyimpan data input ketiga yang telah di isi user ke dalam variabel tiga
 fmt.Scanln(&tiga)

 // Program akan memunculkan output dari variabel yang telah di isi tadi
 fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)

 // Proses pertukaran posisi
 temp = satu // Mengganti "temp" menjadi nilai "satu"
 satu = dua // Mengganti nilai "satu" menjadi nilai "dua"
 dua = tiga // Mengganti nilai "dua" menjadi nilai "tiga"
 tiga = temp // Mengganti nilai "tiga" menjadi nilai "temp" (temp telah berisi nilai "satu")

  // Program akan memunculkan output dari variabel yang telah di tukar posisi tadi
 fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}

//proggram mengganti atau rotasi nilai setiap variabel yang telah di isi oleh user dan mengeluarkan output awal (sebelum rotasi) dan output akhir (setelah dirotasi)