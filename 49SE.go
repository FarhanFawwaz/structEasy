package main

import "fmt"

// Definisikan tipe bentukan untuk data gaji pegawai
type Gaji struct {
	GajiPokok        int
	PajakPenghasilan int
	Asuransi         int
}

func main() {
	// Input data gaji pegawai
	var g Gaji
	fmt.Scan(&g.GajiPokok, &g.PajakPenghasilan, &g.Asuransi)

	// Hitung total gaji
	totalGaji := hitungTotalGaji(g)

	// Tampilkan total gaji
	fmt.Println(totalGaji)
}

// Fungsi untuk menghitung total gaji pegawai
func hitungTotalGaji(g Gaji) int {
	// Hitung pajak penghasilan
	pajak := float64(g.GajiPokok * g.PajakPenghasilan / 100)

	// Hitung asuransi
	asuransi := float64(g.GajiPokok * g.Asuransi / 100)

	// Hitung total gaji setelah pajak dan asuransi
	totalGaji := g.GajiPokok - int(pajak) - int(asuransi)

	return totalGaji
}
