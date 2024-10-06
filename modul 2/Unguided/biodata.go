package main

import "fmt"

func main() {
	// Membuat 3 variabel yaitu Nama, NIM, Kelas
	var nama, nim, kelas string
    // Meminta user untuk memasukkan Nama 
	fmt.Print("Masukkan Nama Mahasiswa = ")
	// Scan menyimpan data yang telah di isi user ke dalam variabel nama
	fmt.Scan(&nama)
	
    // Meminta user untuk memasukkan NIM 
	fmt.Print("Masukkan NIM Mahasiswa = ")
	// Scan menyimpan data yang telah di isi user ke dalam variabel NIM
	fmt.Scan(&nim)
	
    // Meminta user untuk memasukkan Kelas 
	fmt.Print("Masukkan Kelas Mahasiswa = ")
	// Scan menyimpan data yang telah di isi user ke dalam variabel Kelas
	fmt.Scan(&kelas)
	
	// Semua data yang di isi akan di tampilkan di sini dengan tambahan beberapa kata dan varibael
	fmt.Print("Perkenalkan saya adalah ", nama, ", saya salah satu mahasiswa Prodi S1-IF dari kelas ", kelas, " dengan NIM ", nim,)
}
