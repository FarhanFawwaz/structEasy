package main

import "fmt"

type Parkir struct {
	JenisKendaraan string
	DurasiParkir   int
}

func main() {
	var totalPendapatan int

	for {
		var parkir Parkir

		fmt.Scan(&parkir.JenisKendaraan)

		// Periksa apakah jenis kendaraan valid
		if parkir.JenisKendaraan != "motor" && parkir.JenisKendaraan != "mobil" {
			break
		}

		fmt.Scan(&parkir.DurasiParkir)

		// Periksa apakah durasi parkir valid
		if parkir.DurasiParkir < 0 {
			break
		}

		// Hitung tarif parkir
		tarif := hitungTarif(parkir)
		totalPendapatan += tarif
	}

	// Tampilkan total pendapatan
	fmt.Println(totalPendapatan)
}

func hitungTarif(parkir Parkir) int {
	var tarifPerJam int

	// Tentukan tarif berdasarkan jenis kendaraan
	if parkir.JenisKendaraan == "motor" {
		tarifPerJam = 1000
	} else if parkir.JenisKendaraan == "mobil" {
		tarifPerJam = 5000
	}

	// Hitung total tarif
	totalTarif := tarifPerJam * parkir.DurasiParkir
	return totalTarif
}
