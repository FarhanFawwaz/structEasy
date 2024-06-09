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
	fmt.Println(cetak_pemain_pointMin(p1, p2, p3))
	fmt.Println(cetak_pemain_golMin(p1, p2, p3))
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.point, &p1.selisihgol)
	fmt.Scanln(&p2.nama, &p2.point, &p2.selisihgol)
	fmt.Scanln(&p3.nama, &p3.point, &p3.selisihgol)
}

func cetak_pemain_pointMin(p1, p2, p3 bola) string {
	var min bola
	min = p1
	if p2.point < min.point {
		min = p2
	}
	if p3.point < min.point {
		min = p3
	}
	return min.nama
}
func cetak_pemain_golMin(p1, p2, p3 bola) string {
	var min bola
	min = p1
	if p2.selisihgol < min.selisihgol {
		min = p2
	}
	if p3.selisihgol < min.selisihgol {
		min = p3
	}
	return min.nama
}
