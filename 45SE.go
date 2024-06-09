package main

import (
	"fmt"
	"math"
)

// Definisikan tipe bentukan untuk review series
type SeriesReview struct {
	NamaSeries    string
	JumlahEpisode int
	DurasiPerEps  int
	DurasiTotal   int
}

func main() {
	// Input data series
	var series SeriesReview

	fmt.Scan(&series.NamaSeries)
	fmt.Scan(&series.JumlahEpisode)

	fmt.Scan(&series.DurasiPerEps)

	// Hitung durasi total series
	series.DurasiTotal = series.JumlahEpisode * series.DurasiPerEps

	// Input jam yang bisa ditonton dalam sehari
	var jamPerHari int

	fmt.Scan(&jamPerHari)

	// Hitung jumlah hari untuk menonton series
	jumlahHari := hitungJumlahHari(series.DurasiTotal, jamPerHari)

	// Output hasil
	fmt.Print(jumlahHari, " hari")
}

// Fungsi untuk menghitung jumlah hari
func hitungJumlahHari(durasiTotal, jamPerHari int) int {
	return int(math.Ceil(float64(durasiTotal) / float64(jamPerHari*60)))
}
