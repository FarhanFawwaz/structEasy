package main

import "fmt"

type HasilGenap struct {
	BilanganTerakhir int
	JumlahTotal      int
}

func jumlahBilanganGenap(n int) HasilGenap {
	var total, bilanganTerakhir int

	// Menghitung total dan bilangan genap terakhir
	for i := 1; i <= n; i++ {
		bilangan := 2 * i
		total += bilangan
		bilanganTerakhir = bilangan
	}

	return HasilGenap{
		BilanganTerakhir: bilanganTerakhir,
		JumlahTotal:      total,
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	hasil := jumlahBilanganGenap(n)

	fmt.Println(hasil.BilanganTerakhir, hasil.JumlahTotal)
}
