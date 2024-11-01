package main

import "fmt"
func main() {
 var bilangan int
 var teks string
 fmt.Scan(&bilangan)
 teks = "bukan"
 if bilangan % 2 == 0 && bilangan < 0 {
 teks = "genap negatif"
 }
 fmt.Println(teks)
}
