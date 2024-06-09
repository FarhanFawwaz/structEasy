package main

import (
	"fmt"
)

// Definisikan tipe bentukan koordinat
type koordinat struct {
	X, Y int
}

func main() {
	// Membaca koordinat harta karun
	var hartaKarun koordinat
	fmt.Scan(&hartaKarun.X, &hartaKarun.Y)

	// Percobaan penggalian
	for i := 0; i < 3; i++ {
		// Membaca tebakan koordinat
		var tebakan koordinat
		fmt.Scan(&tebakan.X, &tebakan.Y)

		// Memeriksa apakah tebakan benar
		if tebakan.X == hartaKarun.X && tebakan.Y == hartaKarun.Y {
			fmt.Println("Yay Menang")
			return
		} else {
			fmt.Println("Coba Lagi")
		}
	}

	// Jika semua percobaan habis dan harta karun belum ditemukan
}
