package main

import "fmt"

type HasilPerkalian struct {
	TotalProduk    float64
	JumlahBilangan int
}

func perkalian() HasilPerkalian {
	var input float64
	hasil := HasilPerkalian{TotalProduk: 1, JumlahBilangan: 0}

	for {
		fmt.Scanln(&input)

		if input == 0 {
			break
		}

		hasil.TotalProduk *= (1 / input)
		hasil.JumlahBilangan++
	}

	return hasil
}

func main() {
	hasilPerkalian := perkalian()

	fmt.Println(hasilPerkalian.TotalProduk)
	fmt.Println(hasilPerkalian.JumlahBilangan)
}
