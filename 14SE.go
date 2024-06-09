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
	var keliling float64
	bacaData(&A, &B, &C)

	keliling = hitungKeliling(A, B, C)
	fmt.Println(keliling)
}

func bacaData(A, B, C *Titik) {
	fmt.Scanln(&A.X, &A.Y)
	fmt.Scanln(&B.X, &B.Y)
	fmt.Scanln(&C.X, &C.Y)
}

func hitungKeliling(A, B, C Titik) float64 {
	// Menggunakan rumus jarak Euclidean antara dua titik dalam ruang dua dimensi
	AB := jarak(A, B)
	BC := jarak(B, C)
	AC := jarak(A, C)

	// Menghitung panjang keliling segitiga
	keliling := AB + BC + AC
	return keliling
}

func jarak(A, B Titik) float64 {
	// Menggunakan rumus jarak Euclidean antara dua titik dalam ruang dua dimensi
	return math.Sqrt(math.Pow(B.X-A.X, 2) + math.Pow(B.Y-A.Y, 2))
}
