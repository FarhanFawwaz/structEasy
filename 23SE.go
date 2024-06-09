package main

import "fmt"

type Waktu struct {
	Jam   int
	Menit int
	Detik int
}

func hitungDurasiMasukKeluar(masuk, keluar Waktu) Waktu {
	// Konversi semua waktu ke detik
	totalDetikMasuk := masuk.Jam*3600 + masuk.Menit*60 + masuk.Detik
	totalDetikKeluar := keluar.Jam*3600 + keluar.Menit*60 + keluar.Detik

	// Hitung selisih waktu dalam detik
	selisihDetik := totalDetikKeluar - totalDetikMasuk

	// Konversi kembali ke jam, menit, dan detik
	jam := selisihDetik / 3600
	selisihDetik %= 3600
	menit := selisihDetik / 60
	detik := selisihDetik % 60

	return Waktu{Jam: jam, Menit: menit, Detik: detik}
}

func main() {
	var jamMasuk, menitMasuk, detikMasuk int
	var jamKeluar, menitKeluar, detikKeluar int

	// Baca input waktu masuk
	fmt.Scan(&jamMasuk, &menitMasuk, &detikMasuk)
	waktuMasuk := Waktu{Jam: jamMasuk, Menit: menitMasuk, Detik: detikMasuk}

	// Baca input waktu keluar
	fmt.Scan(&jamKeluar, &menitKeluar, &detikKeluar)
	waktuKeluar := Waktu{Jam: jamKeluar, Menit: menitKeluar, Detik: detikKeluar}

	// Hitung durasi parkir
	durasiParkir := hitungDurasiMasukKeluar(waktuMasuk, waktuKeluar)

	// Cetak hasil
	fmt.Printf("%d %d %d\n", durasiParkir.Jam, durasiParkir.Menit, durasiParkir.Detik)
}
