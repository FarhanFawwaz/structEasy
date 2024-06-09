package main

import "fmt"

const pi float64 = 3.14

// Definisikan tipe bentukan bola
type Bola struct {
	Nama   string
	Jari   int
	Volume float64
}

func main() {
	var maxVolume float64
	var namaBolaTerbesar string

	// Inisialisasi volume terbesar dengan nilai 0
	maxVolume = 0

	// Input data bola dan hitung volume
	for i := 0; i < 5; i++ {
		var bola Bola

		fmt.Scan(&bola.Nama, &bola.Jari)

		// Hitung volume bola
		bola.Volume = hitungVolumeBola(bola.Jari)

		// Periksa apakah volume bola lebih besar dari volume terbesar sejauh ini
		if bola.Volume > maxVolume {
			maxVolume = bola.Volume
			namaBolaTerbesar = bola.Nama
		}
	}

	// Tampilkan nama dan volume bola terbesar
	fmt.Print("", namaBolaTerbesar, " ", maxVolume)
}

// Fungsi untuk menghitung volume bola
func hitungVolumeBola(jari int) float64 {
	return (4.0 / 3.0) * pi * float64(jari*jari*jari)
}
