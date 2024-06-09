package main

import (
	"fmt"
)

// Deklarasikan tipe bentukan untuk menampung data langkah
type Langkah struct {
	BuahCatur   string
	PetakAsal   string
	PetakTujuan string
	Nilai       int
}

func main() {
	var n int = 3 // Jumlah langkah yang akan diinput, bisa diubah sesuai kebutuhan

	// Membuat slice untuk menampung data langkah
	langkahs := make([]Langkah, n)

	// Membaca input langkah-langkah dari pengguna
	for i := 0; i < n; i++ {
		fmt.Scan(&langkahs[i].BuahCatur, &langkahs[i].PetakAsal, &langkahs[i].PetakTujuan, &langkahs[i].Nilai)
	}

	// Cari langkah dengan nilai tertinggi
	maxIndex := 0
	for i := 1; i < len(langkahs); i++ {
		if langkahs[i].Nilai > langkahs[maxIndex].Nilai {
			maxIndex = i
		}
	}

	// Dapatkan langkah dengan nilai tertinggi
	langkahTertinggi := langkahs[maxIndex]
	hasil := fmt.Sprintf("%s %s-%s", langkahTertinggi.BuahCatur, langkahTertinggi.PetakAsal, langkahTertinggi.PetakTujuan)

	// Cetak hasil
	fmt.Println(hasil)
}
