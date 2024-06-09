package main

import (
	"fmt"
	"math"
)

type Mahasiswa struct {
	Nama string
	NIM  int
	IPK  float64
}

func main() {
	// Membuat slice untuk menampung data mahasiswa
	mahasiswas := make([]Mahasiswa, 3)

	// Membaca input data mahasiswa
	for i := 0; i < 3; i++ {
		fmt.Scan(&mahasiswas[i].Nama, &mahasiswas[i].NIM, &mahasiswas[i].IPK)
	}

	// Menghitung rata-rata IPK
	var totalIPK float64
	maxIPK := mahasiswas[0].IPK
	namaMaxIPK := mahasiswas[0].Nama

	for _, mhs := range mahasiswas {
		totalIPK += mhs.IPK
		if mhs.IPK > maxIPK {
			maxIPK = mhs.IPK
			namaMaxIPK = mhs.Nama
		}
	}

	rataRataIPK := totalIPK / 3

	// Cetak hasil dengan pembulatan 2 digit di belakang koma
	fmt.Printf("%.2f\n", math.Round(rataRataIPK*100)/100)
	fmt.Println(namaMaxIPK)
}
