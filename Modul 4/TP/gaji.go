package main

import "fmt"

func main () {
    var jamKerjaMingguan, upahPerJam, gajiMingguan, jamLembur, gajiLembur float64

    // Meminta input jumlah jam kerja dalam seminggu 
    fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
    fmt.Scanln(&jamKerjaMingguan)
    // Meminta input upah perjam
    fmt.Print("Masukkan upah per jam: ")
    fmt.Scanln(&upahPerJam)

    // Menghitung gaji mingguan
    // if kita gunakan untuk memeriksa apakah jam kerja karyawan lebih dari 40 hari
    if jamKerjaMingguan > 40 {
        // Menghitung jumlah jam lembur dengan cara jam kerja mingguan dikurangi 40
        jamLembur = jamKerjaMingguan - 40
        // menghitung gaji lembur denfan cara jam lembur dikali upah kerja dan di kali lagi dengan bonus lembur yang di bayar 1.5 kali lipat
        gajiLembur = jamLembur * upahPerJam * 1.5
        // Menghitung gaji mingguan dengan cara upah perjam di kali dengan 40 jam biasa lalu di tambah dengan gaji lembur
        gajiMingguan = (40 * upahPerJam) + gajiLembur
        // else berguna untuk jika kondisi if tidak terpenuhi atau jam kerja karyawan di bawah 40 hari
    } else {
        // jika tidak memenuhi syarat dari if yaitu kurang dari 40 jam kerja maka cukup mengalilan jam kerja mingguan dengan upah perjam
        gajiMingguan = jamKerjaMingguan * upahPerJam
    }

    // Menghitung total gaji bulanan dengan cara gaji mingguan di kali 4 yaitu jumlah minggu dalam 1 bulan
    totalGajiBulanan := gajiMingguan * 4

    // Menampilkan hasil gaji bulanan
    fmt.Print("Total gaji bulanan adalah: Rp.", totalGajiBulanan)
}
