package main

import "fmt"

type bola struct {
	nama       string
	point      int
	selisihGol int
}

func main() {
	var p1, p2, p3 bola
	bacaData(&p1, &p2, &p3)
	cetakRataPoint(p1, p2, p3)
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.point, &p1.selisihGol)
	fmt.Scanln(&p2.nama, &p2.point, &p2.selisihGol)
	fmt.Scanln(&p3.nama, &p3.point, &p3.selisihGol)
}

func cetakRataPoint(p1, p2, p3 bola) {
	rataPoint := rataRataPoint(p1, p2, p3)
	fmt.Println(rataPoint)
}

func rataRataPoint(p1, p2, p3 bola) float64 {
	rataPoint := float64(p1.point + p2.point + p3.point)
	return rataPoint / 3
}
