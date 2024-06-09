package main

import "fmt"

type bola struct {
	nama   string
	gol    int
	assist int
}

func main() {
	var p1, p2, p3 bola
	bacaData(&p1, &p2, &p3)
	cetakRataGol(p1, p2, p3)
	cetakRataAssist(p1, p2, p3)
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.gol, &p1.assist)
	fmt.Scanln(&p2.nama, &p2.gol, &p2.assist)
	fmt.Scanln(&p3.nama, &p3.gol, &p3.assist)
}

func cetakRataGol(p1, p2, p3 bola) {
	rataGol := rataRataGol(p1, p2, p3)
	fmt.Println(rataGol)
}

func rataRataGol(p1, p2, p3 bola) float64 {
	totalGol := float64(p1.gol + p2.gol + p3.gol)
	return totalGol / 3
}

func cetakRataAssist(p1, p2, p3 bola) {
	rataAssist := rataRataAssist(p1, p2, p3)
	fmt.Println(rataAssist)
}

func rataRataAssist(p1, p2, p3 bola) float64 {
	totalAssist := float64(p1.assist + p2.assist + p3.assist)
	return totalAssist / 3
}
