package main

import (
	"fmt"
	"math"
)

type Titik struct {
	X float64
	Y float64
}

type Garis struct {
	Titik1 Titik
	Titik2 Titik
}

func panjangGaris(garis Garis) float64 {
	dx := garis.Titik2.X - garis.Titik1.X
	dy := garis.Titik2.Y - garis.Titik1.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func garisSamaPanjang(garis1, garis2 Garis) bool {
	panjang1 := panjangGaris(garis1)
	panjang2 := panjangGaris(garis2)
	return math.Abs(panjang1-panjang2) < 1e-9
}

func main() {
	var titikA, titikB, titikC, titikD Titik

	fmt.Scanln(&titikA.X, &titikA.Y, &titikB.X, &titikB.Y, &titikC.X, &titikC.Y, &titikD.X, &titikD.Y)

	garis1 := Garis{Titik1: titikA, Titik2: titikB}
	garis2 := Garis{Titik1: titikC, Titik2: titikD}

	samaPanjang := garisSamaPanjang(garis1, garis2)

	fmt.Println(samaPanjang)
}
