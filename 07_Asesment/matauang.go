package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
	fmt.Scan(&qirat)

	dinar := qirat / (10 * 10 * 6)
	sisa := qirat % (10 * 10 * 6)

	dirham := sisa / (10 * 6)
	sisa = sisa % (10 * 6)

	fals := sisa / 6
	sisa = sisa % 6

	qirat = sisa

	fmt.Println(dinar, dirham, fals, qirat)
}