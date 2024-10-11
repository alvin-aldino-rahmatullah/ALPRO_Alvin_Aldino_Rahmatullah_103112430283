package main

import "fmt"

func main() { 

  var BelanjaAwal, BelanjaAkhir, diskon float64
  fmt.Print("Masukkan Total Belanja: " )
  fmt.Scan(&BelanjaAwal)
  fmt.Print("Jumlah Diskon: ")
  fmt.Scan(&diskon)
  
  BelanjaAkhir = BelanjaAwal  - (BelanjaAwal*diskon/100)
   
  fmt.Print("Total Belanja Akhir: ", BelanjaAkhir)

}