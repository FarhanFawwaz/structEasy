package main

import "fmt"

type HasilGanjil struct {
	BilanganTerakhir int
	JumlahTotal      int
	RataRata         float64
}

func jumlahBilanganGanjil(n int) HasilGanjil {
	var bilangan, jumlah int
	for i := 1; i <= n; i++ {
		bilangan = 2*i - 1 // Bilangan ganjil adalah 2n - 1
		jumlah += bilangan
	}
	rataRata := float64(jumlah) / float64(n)
	hg := HasilGanjil{
		BilanganTerakhir: bilangan,
		JumlahTotal:      jumlah,
		RataRata:         rataRata,
	}
	return hg
}

func main() {
	var n int
	fmt.Scan(&n)

	hasil := jumlahBilanganGanjil(n)

	fmt.Println(hasil.BilanganTerakhir, hasil.JumlahTotal, hasil.RataRata)
}
