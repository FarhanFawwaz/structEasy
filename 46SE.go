package main

import "fmt"

// Definisikan tipe bentukan untuk pesanan sushi berdasarkan warna
type PesananSushi struct {
	Hitam  int
	Biru   int
	Kuning int
	Merah  int
}

func main() {
	// Input jumlah pesanan sushi dari tiga pelanggan
	var pelanggan [3]PesananSushi

	for i := 0; i < 3; i++ {

		fmt.Scan(&pelanggan[i].Hitam, &pelanggan[i].Biru, &pelanggan[i].Kuning, &pelanggan[i].Merah)
	}

	// Hitung total biaya pesanan dari tiga pelanggan
	totalBiaya := hitungTotalBiaya(pelanggan)

	// Tampilkan total biaya keseluruhan
	fmt.Print(totalBiaya)
}

// Fungsi untuk menghitung total biaya pesanan
func hitungTotalBiaya(pelanggan [3]PesananSushi) int {
	var total int
	for i := 0; i < 3; i++ {
		// Hitung biaya pesanan untuk setiap pelanggan
		biayaPelanggan := (pelanggan[i].Hitam * 50000) +
			(pelanggan[i].Biru * 20000) +
			(pelanggan[i].Kuning * 10000) +
			(pelanggan[i].Merah * 5000)
		// Tambahkan biaya pesanan pelanggan ke total biaya
		total += biayaPelanggan
	}
	// Tambahkan pajak 10%
	total += total * 10 / 100
	return total
}
