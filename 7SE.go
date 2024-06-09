package main

import "fmt"

type bola struct {
	nama       string
	point      int
	selisihgol int
}

func main() {
	var p1, p2, p3 bola
	bacaData(&p1, &p2, &p3)
	fmt.Println(cetak_pemain_pointMaks(p1, p2, p3))
	fmt.Println(cetak_pemain_golMaks(p1, p2, p3))
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.point, &p1.selisihgol)
	fmt.Scanln(&p2.nama, &p2.point, &p2.selisihgol)
	fmt.Scanln(&p3.nama, &p3.point, &p3.selisihgol)
}

func cetak_pemain_pointMaks(p1, p2, p3 bola) string {
	var max bola
	max = p1
	if p2.point > max.point {
		max = p2
	}
	if p3.point > max.point {
		max = p3
	}
	return max.nama
}
func cetak_pemain_golMaks(p1, p2, p3 bola) string {
	var max bola
	max = p1
	if p2.selisihgol > max.selisihgol {
		max = p2
	}
	if p3.selisihgol > max.selisihgol {
		max = p3
	}
	return max.nama
}
