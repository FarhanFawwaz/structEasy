package main

import "fmt"

// Tipe bentukan untuk menyatakan data penanggalan
type Tanggal struct {
	Hari  int
	Bulan int
	Tahun int
}

// Fungsi untuk menghitung jumlah hari dari awal tahun sampai tanggal tertentu
func jumlahHari(t Tanggal) int {
	return t.Hari + (t.Bulan-1)*30 + (t.Tahun-1)*360
}

func main() {
	var tglPinjam, blnPinjam, thnPinjam int
	var tglKembali, blnKembali, thnKembali int

	// Membaca input tanggal peminjaman
	fmt.Scan(&tglPinjam, &blnPinjam, &thnPinjam)
	tanggalPinjam := Tanggal{Hari: tglPinjam, Bulan: blnPinjam, Tahun: thnPinjam}

	// Membaca input tanggal pengembalian
	fmt.Scan(&tglKembali, &blnKembali, &thnKembali)
	tanggalKembali := Tanggal{Hari: tglKembali, Bulan: blnKembali, Tahun: thnKembali}

	// Menghitung jumlah hari peminjaman
	hariPinjam := jumlahHari(tanggalPinjam)
	hariKembali := jumlahHari(tanggalKembali)
	lamaPinjaman := hariKembali - hariPinjam

	// Mencetak hasil lama hari peminjaman
	fmt.Println(lamaPinjaman)
}
