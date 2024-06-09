package main

import "fmt"

type pasukan struct {
	nama               string
	artileri, kavaleri int
}

func main() {
	var p1, p2, p3 pasukan
	bacaData(&p1, &p2, &p3)
	fmt.Println(cetakGabungan(p1, p2, p3))
}
func bacaData(p1, p2, p3 *pasukan) {
	fmt.Scanln(&p1.nama, &p1.artileri, &p1.kavaleri)
	fmt.Scanln(&p2.nama, &p2.artileri, &p2.kavaleri)
	fmt.Scanln(&p3.nama, &p3.artileri, &p3.kavaleri)
}
func cetakGabungan(p1, p2, p3 pasukan) string {
	var totalP1, totalP2, totalP3 int
	totalP1 = p1.artileri + p1.kavaleri
	totalP2 = p2.artileri + p2.kavaleri
	totalP3 = p3.artileri + p3.kavaleri

	if totalP1 >= totalP2 && totalP1 >= totalP3 {
		return p1.nama
	} else if totalP2 >= totalP1 && totalP2 >= totalP3 {
		return p2.nama
	} else {
		return p3.nama
	}
}
