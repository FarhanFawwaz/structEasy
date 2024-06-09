package main

import "fmt"

// Definisikan tipe bentukan untuk tim balap formula one
type TimBalap struct {
	Tim    string
	Driver string
	Poin   int
}

func main() {
	var timA, timB, timC TimBalap

	// Input data tim
	fmt.Scan(&timA.Tim, &timA.Driver, &timA.Poin)
	fmt.Scan(&timB.Tim, &timB.Driver, &timB.Poin)
	fmt.Scan(&timC.Tim, &timC.Driver, &timC.Poin)

	// Proses input poin balap
	var input string
	var poin int
	for {
		fmt.Scan(&input)
		if input == "SELESAI" {
			break
		}
		fmt.Scan(&poin)
		switch input {
		case timA.Driver:
			timA.Poin += poin
		case timB.Driver:
			timB.Poin += poin
		case timC.Driver:
			timC.Poin += poin
		}
	}

	// Tentukan tim dan driver pemenang
	var pemenang TimBalap
	if timA.Poin > timB.Poin && timA.Poin > timC.Poin {
		pemenang = timA
	} else if timB.Poin > timA.Poin && timB.Poin > timC.Poin {
		pemenang = timB
	} else {
		pemenang = timC
	}

	// Tampilkan hasil
	fmt.Println(pemenang.Tim, pemenang.Driver)
}
