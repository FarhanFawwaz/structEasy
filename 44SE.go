package main

import "fmt"

type Order struct {
	NomorMeja  string
	Makanan    string
	Minuman    string
	TotalHarga int
}

func main() {
	var totalHargaSemua, totalPesanan, count int
	var order1, order2 Order // Variabel order1 dan order2 untuk menyimpan data order

	// Harga menu
	hargaMenu := map[string]int{
		"Dimsum": 20000,
		"Bakwan": 12000,
		"Es_Teh": 7000,
		"Air":    3000,
	}

	// Input data order
	for i := 0; i < 2; i++ {
		var order Order

		fmt.Scan(&order.NomorMeja)
		fmt.Scan(&order.Makanan)
		fmt.Scan(&order.Minuman)

		// Hitung total harga pesanan
		order.TotalHarga = hargaMenu[order.Makanan] + hargaMenu[order.Minuman]
		totalHargaSemua += order.TotalHarga
		totalPesanan += order.TotalHarga
		count++

		if i == 0 {
			order1 = order // Simpan order pertama
		} else {
			order2 = order // Simpan order kedua
		}
	}

	// Hitung rata-rata total pesanan

	// Tentukan nomor meja dengan prioritas berdasarkan total harga
	nomorMejaPrioritas := orderPrioritas(order1, order2)

	// Output nomor meja dengan prioritas utama
	fmt.Printf("Nomor Meja: %s\n", nomorMejaPrioritas)
	if nomorMejaPrioritas > "1" {
		fmt.Printf("Nomor Meja: %s\n", nomorMejaPrioritas)
	}
	// Output total pesanan hari ini
	fmt.Printf("Total order hari ini: Rp %d\n", totalPesanan)

}

func orderPrioritas(order1, order2 Order) string {
	if order1.TotalHarga > order2.TotalHarga {
		return order1.NomorMeja
	} else if order1.TotalHarga < order2.TotalHarga {
		return order2.NomorMeja
	}
	return order1.NomorMeja + " " + order2.NomorMeja
}
