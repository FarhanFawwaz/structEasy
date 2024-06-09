package main

import (
	"fmt"
)

// Definisikan tipe bentukan untuk data tubuh
type dataTubuh struct {
	Tinggi float64
	Berat  float64
	BMI    float64
}

func main() {
	var tinggi, berat float64
	fmt.Scan(&tinggi, &berat)

	// Menghitung BMI
	bmi := hitungBMI(tinggi, berat)

	// Menentukan kategori BMI
	kategori := tentukanKategoriBMI(bmi)

	// Menampilkan hasil
	fmt.Printf("BMI: %.2f\n", bmi)
	fmt.Println("Kategori:", kategori)
}

// Fungsi untuk menghitung BMI
func hitungBMI(tinggi, berat float64) float64 {
	return berat / (tinggi * tinggi)
}

// Fungsi untuk menentukan kategori BMI
func tentukanKategoriBMI(bmi float64) string {
	if bmi < 18.5 {
		return "Underweight"
	} else if bmi >= 18.5 && bmi < 24.9 {
		return "Normal Weight"
	} else if bmi >= 24.9 && bmi < 29.9 {
		return "Overweight"
	} else {
		return "Obesity"
	}
}
