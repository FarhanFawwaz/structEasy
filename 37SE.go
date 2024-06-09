package main

import "fmt"

// Definisikan tipe bentukan balok
type Balok struct {
	Panjang int
	Lebar   int
	Tinggi  int
	Ukuran  Geometri
}

// Definisikan tipe bentukan geometri
type Geometri struct {
	LuasAlas int
	Volume   int
}

func main() {
	var balok Balok

	// Input ukuran panjang, lebar, dan tinggi balok
	fmt.Scan(&balok.Panjang, &balok.Lebar, &balok.Tinggi)

	// Hitung luas alas dan volume balok
	hitungBalok(&balok)

	// Tampilkan hasil
	fmt.Print(balok.Ukuran.LuasAlas, balok.Ukuran.Volume)

}

// Fungsi untuk menghitung luas alas dan volume balok
func hitungBalok(balok *Balok) {
	balok.Ukuran.LuasAlas = balok.Panjang * balok.Lebar
	balok.Ukuran.Volume = balok.Panjang * balok.Lebar * balok.Tinggi
}
