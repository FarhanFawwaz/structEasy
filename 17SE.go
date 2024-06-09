package main

import "fmt"

type DataHarmonis struct {
	JumlahKebalikan float64
	JumlahBilangan  int
}

func rataHarmonis(n int) float64 {
	var sumHarmonic float64
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanln(&num)
		sumHarmonic += 1.0 / float64(num)
	}
	return float64(n) / sumHarmonic
}

func main() {
	var n int
	fmt.Scanln(&n)

	rata := rataHarmonis(n)

	fmt.Printf("%.2f\n", rata)
}
