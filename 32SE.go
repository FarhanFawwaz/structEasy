package main

import "fmt"

type Atlet struct {
	Asal  string
	Waktu float64
}

func main() {
	var atlet1, atlet2, atlet3 Atlet

	// Membaca input data atlet
	fmt.Scan(&atlet1.Asal, &atlet1.Waktu)
	fmt.Scan(&atlet2.Asal, &atlet2.Waktu)
	fmt.Scan(&atlet3.Asal, &atlet3.Waktu)

	pemenang(&atlet1, &atlet2, &atlet3)
}

func pemenang(atlet1, atlet2, atlet3 *Atlet) {
	// Memeriksa apakah waktu ketiga atlet sama
	if atlet1.Waktu == atlet2.Waktu && atlet2.Waktu == atlet3.Waktu {
		fmt.Println("seri")
		return
	}

	// Mencari pemenang berdasarkan waktu tercepat
	var pemenang string
	if atlet1.Waktu < atlet2.Waktu && atlet1.Waktu < atlet3.Waktu {
		pemenang = atlet1.Asal
	} else if atlet2.Waktu < atlet1.Waktu && atlet2.Waktu < atlet3.Waktu {
		pemenang = atlet2.Asal
	} else {
		pemenang = atlet3.Asal
	}

	fmt.Println("Atlet asal", pemenang, "menang")
}
