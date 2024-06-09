package main

import "fmt"

type KasetGame struct {
	Judul    string
	Harga    int
	Genre    string
	Preorder bool
}

func main() {
	var kaset1, kaset2, kaset3 KasetGame

	// Input data kaset game
	fmt.Scan(&kaset1.Judul, &kaset1.Harga, &kaset1.Genre, &kaset1.Preorder)
	fmt.Scan(&kaset2.Judul, &kaset2.Harga, &kaset2.Genre, &kaset2.Preorder)
	fmt.Scan(&kaset3.Judul, &kaset3.Harga, &kaset3.Genre, &kaset3.Preorder)

	// Hitung dan tampilkan harga kaset game setelah diskon
	fmt.Println(hitungDiskon(kaset1))
	fmt.Println(hitungDiskon(kaset2))
	fmt.Println(hitungDiskon(kaset3))
}

func hitungDiskon(k KasetGame) (judul, genre string, harga int) {
	// Inisialisasi diskon awal
	diskon := 0

	// Cek diskon untuk pre-order
	if k.Preorder {
		diskon += 5
	}

	// Cek diskon tambahan untuk pre-order dengan harga minimal 500.000
	if k.Preorder && k.Harga >= 500000 {
		diskon += 10
	}

	// Cek diskon untuk genre Adventure dan Action
	if k.Genre == "Adventure" || k.Genre == "Action" {
		diskon += 10
	}

	// Batasi diskon maksimal menjadi 15%
	if diskon > 15 {
		diskon = 15
	}

	// Hitung harga setelah diskon
	hargaDiskon := k.Harga - (k.Harga * diskon / 100)

	return k.Judul, k.Genre, hargaDiskon
}
