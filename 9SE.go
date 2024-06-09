package main

import "fmt"

type pasukan struct {
	nama     string
	artileri int
	kavaleri int
}

func main() {
	var p1, p2, p3 pasukan
	bacaData(&p1, &p2, &p3)
	cetakJumlahArtileri(p1, p2, p3)
	cetakJumlahKavaleri(p1, p2, p3)
}

func bacaData(p1, p2, p3 *pasukan) {
	fmt.Scanln(&p1.nama, &p1.artileri, &p1.kavaleri)
	fmt.Scanln(&p2.nama, &p2.artileri, &p2.kavaleri)
	fmt.Scanln(&p3.nama, &p3.artileri, &p3.kavaleri)
}

func cetakJumlahArtileri(p1, p2, p3 pasukan) {
	jumlahArtileri := totArtileri(p1, p2, p3)
	fmt.Print(jumlahArtileri)
}

func totArtileri(p1, p2, p3 pasukan) int {
	totA := (p1.artileri + p2.artileri + p3.artileri)
	return totA
}

func cetakJumlahKavaleri(p1, p2, p3 pasukan) {
	jumlahKavileri := totKavileri(p1, p2, p3)
	fmt.Printf(" %d", jumlahKavileri)
}

func totKavileri(p1, p2, p3 pasukan) int {
	totK := (p1.kavaleri + p2.kavaleri + p3.kavaleri)
	return totK
}
