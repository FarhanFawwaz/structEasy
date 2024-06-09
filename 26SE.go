package main

import "fmt"

// Tipe bentukan untuk menyatakan data waktu
type Waktu struct {
	Jam   int
	Menit int
}

// Tipe bentukan untuk menyatakan data jadwal kereta api
type JadwalKereta struct {
	NomorKA    string
	KotaAsal   string
	Berangkat  Waktu
	KotaTujuan string
	Tiba       Waktu
}

// Fungsi untuk menghitung durasi perjalanan dalam menit
func hitungDurasi(berangkat, tiba Waktu) int {
	menitBerangkat := berangkat.Jam*60 + berangkat.Menit
	menitTiba := tiba.Jam*60 + tiba.Menit
	return menitTiba - menitBerangkat
}

func main() {
	// Membuat slice untuk menampung data jadwal kereta
	jadwals := make([]JadwalKereta, 3)

	// Membaca input data jadwal kereta
	for i := 0; i < 3; i++ {
		var nomorKA, kotaAsal, kotaTujuan string
		var jamBerangkat, menitBerangkat, jamTiba, menitTiba int
		fmt.Scan(&nomorKA, &kotaAsal, &jamBerangkat, &menitBerangkat, &kotaTujuan, &jamTiba, &menitTiba)
		jadwals[i] = JadwalKereta{
			NomorKA:    nomorKA,
			KotaAsal:   kotaAsal,
			Berangkat:  Waktu{Jam: jamBerangkat, Menit: menitBerangkat},
			KotaTujuan: kotaTujuan,
			Tiba:       Waktu{Jam: jamTiba, Menit: menitTiba},
		}
	}

	// Cari jadwal kereta dengan durasi terlama
	maxDurasi := hitungDurasi(jadwals[0].Berangkat, jadwals[0].Tiba)
	nomorKATerlama := jadwals[0].NomorKA

	for _, jadwal := range jadwals[1:] {
		durasi := hitungDurasi(jadwal.Berangkat, jadwal.Tiba)
		if durasi > maxDurasi {
			maxDurasi = durasi
			nomorKATerlama = jadwal.NomorKA
		}
	}

	// Cetak hasil nomor KA dengan durasi terlama
	fmt.Println(nomorKATerlama)
}
