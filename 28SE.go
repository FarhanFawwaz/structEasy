package main

import "fmt"

type HasilBaca struct {
	BilanganTerakhir int
	JumlahBacaan     int
}

func hitung() HasilBaca {
	var bilangan, sebelumnya int
	var hasil HasilBaca
	lanjut := true

	for hasil.JumlahBacaan = 0; hasil.JumlahBacaan < 10 && lanjut; hasil.JumlahBacaan++ {
		fmt.Scan(&bilangan)
		if hasil.JumlahBacaan > 0 && bilangan > sebelumnya {
			lanjut = false
		} else {
			sebelumnya = bilangan
			hasil.BilanganTerakhir = bilangan
		}
	}

	return hasil
}

func main() {
	hasil := hitung()
	fmt.Println(hasil.BilanganTerakhir, hasil.JumlahBacaan)
}
