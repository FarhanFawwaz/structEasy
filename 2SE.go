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
	cetakJumlahGol(p1, p2, p3)
	cetakJumlahAssist(p1, p2, p3)
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.gol, &p1.assist)
	fmt.Scanln(&p2.nama, &p2.gol, &p2.assist)
	fmt.Scanln(&p3.nama, &p3.gol, &p3.assist)
}

func cetakJumlahGol(p1, p2, p3 bola) {
	jumlahGol := jumlahDariGol(p1, p2, p3)
	fmt.Println(jumlahGol)
}

func jumlahDariGol(p1, p2, p3 bola) int {
	Gol := (p1.gol + p2.gol + p3.gol)
	return Gol
}

func cetakJumlahAssist(p1, p2, p3 bola) {
	jumlahAssist := jumlahDariAssist(p1, p2, p3)
	fmt.Println(jumlahAssist)
}

func jumlahDariAssist(p1, p2, p3 bola) int {
	assist := (p1.assist + p2.assist + p3.assist)
	return assist
}
