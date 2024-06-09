package main

import (
	"fmt"
	"math"
)

type Titik struct {
	X float64
	Y float64
}

func main() {
	var A, B, C Titik
	bacaData(&A, &B, &C)

	dalamLingkaran := cekDalamLingkaran(A, B, C)
	fmt.Println(dalamLingkaran)
}

func bacaData(A, B, C *Titik) {
	fmt.Scanln(&A.X, &A.Y)
	fmt.Scanln(&B.X, &B.Y)
	fmt.Scanln(&C.X, &C.Y)
}

func cekDalamLingkaran(A, B, C Titik) bool {

	// Menghitung jarak antara titik A dan B sebagai radius lingkaran
	radius := jarak(A, B)

	// Menghitung jarak antara titik A dan C
	jarakAC := jarak(A, C)

	// Jika jarak A ke C lebih kecil atau sama dengan radius lingkaran, maka titik C berada dalam lingkaran
	return jarakAC <= radius
}

func jarak(A, B Titik) float64 {
	// Menggunakan rumus jarak Euclidean antara dua titik dalam ruang dua dimensi
	return math.Sqrt(math.Pow(B.X-A.X, 2) + math.Pow(B.Y-A.Y, 2))
}
