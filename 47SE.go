package main

import "fmt"

// Definisikan tipe bentukan untuk data siswa
type siswa struct {
	Nama        string
	JumlahSuara int
}

func main() {
	// Input nama calon ketua kelas
	var calon1, calon2, calon3 siswa
	fmt.Println("Masukkan nama calon ketua kelas:")
	fmt.Scan(&calon1.Nama, &calon2.Nama, &calon3.Nama)

	// Input jumlah suara untuk setiap calon ketua kelas
	fmt.Println("Masukkan jumlah suara untuk masing-masing calon ketua kelas:")
	inputSuara(&calon1)
	inputSuara(&calon2)
	inputSuara(&calon3)

	// Tentukan pemenang berdasarkan jumlah suara
	pemenang := tentukanPemenang(calon1, calon2, calon3)

	// Tampilkan nama pemenang dan jumlah suaranya
	fmt.Println(pemenang.Nama)
	fmt.Println(pemenang.JumlahSuara)
}

// Fungsi untuk input jumlah suara
func inputSuara(s *siswa) {
	fmt.Scan(&s.JumlahSuara)
}

// Fungsi untuk menentukan pemenang
func tentukanPemenang(calon1, calon2, calon3 siswa) siswa {
	if calon1.JumlahSuara > calon2.JumlahSuara && calon1.JumlahSuara > calon3.JumlahSuara {
		return calon1
	} else if calon2.JumlahSuara > calon1.JumlahSuara && calon2.JumlahSuara > calon3.JumlahSuara {
		return calon2
	} else {
		return calon3
	}
}
