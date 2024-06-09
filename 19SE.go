package main

import (
	"fmt"
	"math"
)

// Mendefinisikan tipe bentukan untuk titik dalam koordinat kartesius 3 dimensi
type Point3D struct {
	x float64
	y float64
	z float64
}

// Fungsi untuk menghitung jarak antara titik dan titik asal (0, 0, 0)
func distanceFromOrigin(p Point3D) float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	var x, y, z float64

	fmt.Scanf("%f %f %f", &x, &y, &z)

	// Membuat titik A
	pointA := Point3D{x: x, y: y, z: z}

	// Menghitung jarak titik A dari titik asal
	distance := distanceFromOrigin(pointA)

	// Menampilkan hasil jarak
	fmt.Print(distance)
}
