package main

import "fmt"

// Definisikan tipe data untuk tiket film
type Tiket struct {
	NamaFilm     string
	Hari         string
	TanggalMerah bool
}

func main() {
	// Input informasi tiket
	var tiket Tiket

	fmt.Scan(&tiket.NamaFilm, &tiket.Hari, &tiket.TanggalMerah)

	// Hitung harga tiket
	harga := hitungHarga(tiket)

	// Tampilkan harga tiket
	fmt.Println(harga)
}

// Fungsi untuk menghitung harga tiket
func hitungHarga(t Tiket) int {
	harga := 50000 // Harga tiket hari kerja

	// Jika hari adalah akhir pekan atau tanggal merah, tambahkan 50%
	if t.Hari == "Sabtu" || t.Hari == "Minggu" || t.TanggalMerah {
		harga += harga * 50 / 100
	}

	return harga
}
