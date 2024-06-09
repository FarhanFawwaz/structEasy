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
	cetakRataSelisihGol(p1, p2, p3)
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.point, &p1.selisihGol)
	fmt.Scanln(&p2.nama, &p2.point, &p2.selisihGol)
	fmt.Scanln(&p3.nama, &p3.point, &p3.selisihGol)
}

func cetakRataSelisihGol(p1, p2, p3 bola) {
	rataSelisihGol := rataRataGol(p1, p2, p3)
	fmt.Println(rataSelisihGol)
}

func rataRataGol(p1, p2, p3 bola) float64 {
	rataGol := float64(p1.selisihGol + p2.selisihGol + p3.selisihGol)
	return rataGol / 3
}
