package main

import "fmt"

// Definisikan tipe bentukan untuk menyimpan data pembeli
type Pembeli struct {
	Member bool
	Harga  int
}

func main() {
	var member bool
	var harga int

	// Membaca input member dan harga belanja
	fmt.Scan(&member, &harga)

	// Menghitung total biaya
	total := hitungBiaya(member, harga)

	// Menampilkan total biaya
	fmt.Println(total)
}

// Fungsi untuk menghitung total biaya
func hitungBiaya(member bool, harga int) float64 {
	var total float64

	// Diskon
	if harga >= 400000 && member {
		total = float64(harga) * 0.8 // Diskon 20%
	} else if harga >= 300000 && member {
		total = float64(harga) * 0.9 // Diskon 10%
	} else if harga >= 300000 {
		total = float64(harga) * 0.8 // Diskon 20%
	} else if harga >= 100000 {
		total = float64(harga) * 0.9 // Diskon 10%
	} else {
		total = float64(harga)
	}

	// Cashback
	if member && harga >= 400000 {
		total -= 50000 // Cashback Rp. 50000
	} else if member && harga >= 200000 {
		total -= 25000 // Cashback Rp. 25000
	}

	return total
}
