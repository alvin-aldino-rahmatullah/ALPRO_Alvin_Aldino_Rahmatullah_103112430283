package main

import (
	"fmt"
	"math"
)

func main() {
    var n int
	fmt.Print("Masukkan jumlah kerucut :")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        var jariJari, tinggi float64

		fmt.Printf("Masukkan jari-jari dan tinggi kerucut : ")
        fmt.Scan(&jariJari, &tinggi)
        
        volume := (1.0 / 3.0) * math.Pi * math.Pow(jariJari, 2) * tinggi

        fmt.Printf("Volume kerucut adalah: %.15f\n", i+1, volume)

    }
}
