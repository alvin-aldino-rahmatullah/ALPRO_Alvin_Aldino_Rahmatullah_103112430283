package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung jarak antara dua titik
func distance(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
    var x1, y1, x2, y2, x3, y3 float64
    
    fmt.Scan(&x1, &y1)
    fmt.Scan(&x2, &y2)
    fmt.Scan(&x3, &y3)
    
    sisiAB := distance(x1, y1, x2, y2) 
    sisiBC := distance(x2, y2, x3, y3) 
    sisiCA := distance(x3, y3, x1, y1) 
    
    maxSisi := math.Max(sisiAB, math.Max(sisiBC, sisiCA))

    fmt.Printf("%.2f", maxSisi)
}
