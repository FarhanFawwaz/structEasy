package main

import (
	"fmt"
	"math"
)

type tpersegi struct {
	x, y float64
}

func main() {
	var A, B, C, D tpersegi
	var keliling float64
	bacaData(&A, &B, &C, &D)
	keliling = hitungkeliling(A, B, C, D)
	fmt.Println(keliling)
}
func bacaData(A, B, C, D *tpersegi) {
	fmt.Scan(&A.x, &A.y)
	fmt.Scan(&B.x, &B.y)
	fmt.Scan(&C.x, &C.y)
	fmt.Scan(&D.x, &D.y)
}
func hitungkeliling(A, B, C, D tpersegi) float64 {
	var k float64
	AB := jarak(A, B)
	BC := jarak(B, C)
	CD := jarak(C, D)
	DA := jarak(D, A)

	k = AB + BC + CD + DA
	return k
}
func jarak(A, B tpersegi) float64 {
	return math.Sqrt(math.Pow(B.x-A.x, 2) + math.Pow(B.y-A.y, 2))
}
