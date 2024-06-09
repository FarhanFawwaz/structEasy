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
	fmt.Println(cetak_pemain_golMaks(p1, p2, p3))
	fmt.Println(cetak_pemain_assistMaks(p1, p2, p3))
}

func bacaData(p1, p2, p3 *bola) {
	fmt.Scanln(&p1.nama, &p1.gol, &p1.assist)
	fmt.Scanln(&p2.nama, &p2.gol, &p2.assist)
	fmt.Scanln(&p3.nama, &p3.gol, &p3.assist)
}

func cetak_pemain_golMaks(p1, p2, p3 bola) string {
	var max bola
	max = p1
	if p2.gol > max.gol {
		max = p2
	}
	if p3.gol > max.gol {
		max = p3
	}
	return max.nama
}
func cetak_pemain_assistMaks(p1, p2, p3 bola) string {
	var max bola
	max = p1
	if p2.assist > max.assist {
		max = p2
	}
	if p3.assist > max.assist {
		max = p3
	}
	return max.nama
}
