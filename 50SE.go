package main

import "fmt"

// Definisikan tipe data untuk menyimpan informasi tentang semangkok mie
type Mie struct {
	TipeMie string
	Pedas   bool
	Topping []string
}

func main() {
	var mie Mie

	// Input data mie

	fmt.Scan(&mie.TipeMie)

	fmt.Scan(&mie.Pedas)

	// Input topping
	for {
		var t string
		fmt.Scan(&t)
		if t == "0" || len(mie.Topping) >= 5 {
			break
		}
		mie.Topping = append(mie.Topping, t)
	}

	// Hitung total harga
	totalHarga := hitungHarga(mie)

	// Tampilkan total harga
	fmt.Println(totalHarga)
}

// Fungsi untuk menghitung total harga mie
func hitungHarga(mie Mie) int {
	// Tentukan harga berdasarkan tipe mie
	var hargaMie int
	switch mie.TipeMie {
	case "kwetiau":
		hargaMie = 8000
	case "bihun":
		hargaMie = 7000
	case "kuning":
		hargaMie = 9000
	}

	// Tambah harga jika mie pedas
	if mie.Pedas {
		hargaMie += 1000
	}

	// Hitung total harga topping
	var hargaTopping int
	for _, t := range mie.Topping {
		switch t {
		case "ayam", "pangsit":
			hargaTopping += 5000
		case "telur", "sayur":
			hargaTopping += 3000
		}
	}

	// Hitung total harga keseluruhan
	totalHarga := hargaMie + hargaTopping

	return totalHarga
}
