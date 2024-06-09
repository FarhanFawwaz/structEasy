package main

import "fmt"

type SkorGame struct {
	TotalA   int
	TotalB   int
	TotalC   int
	min      int
	Pemenang byte
}

var skor SkorGame

func minSkor() {
	if skor.TotalA <= skor.TotalB && skor.TotalA <= skor.TotalC {
		skor.min = skor.TotalA
		skor.Pemenang = 'A'
	} else if skor.TotalB <= skor.TotalA && skor.TotalB <= skor.TotalC {
		skor.min = skor.TotalB
		skor.Pemenang = 'B'
	} else {
		skor.min = skor.TotalC
		skor.Pemenang = 'C'
	}
}

func gameDadu(n int) {
	var d1, d2, d3 int

	for i := 0; i < n; i++ {
		fmt.Scan(&d1, &d2, &d3)
		skor.TotalA += d1
		skor.TotalB += d2
		skor.TotalC += d3
	}
	minSkor()
	fmt.Printf("%c %d\n", skor.Pemenang, skor.min)
}

func main() {
	var n int
	fmt.Scan(&n)

	gameDadu(n)
}
