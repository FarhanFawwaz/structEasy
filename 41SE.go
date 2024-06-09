package main

import "fmt"

type Manajer struct {
	ID              string
	Nama            string
	PengalamanKerja int
}

func main() {
	var bandung, jakarta, bekasi Manajer

	// Input data manajer untuk setiap kota
	fmt.Scan(&bandung.ID, &bandung.Nama, &bandung.PengalamanKerja)
	fmt.Scan(&jakarta.ID, &jakarta.Nama, &jakarta.PengalamanKerja)
	fmt.Scan(&bekasi.ID, &bekasi.Nama, &bekasi.PengalamanKerja)

	// Input data kandidat manajer
	for {
		var kota, id, nama string
		var pengalaman int

		fmt.Scan(&kota)
		if kota == "0" {
			break
		}

		fmt.Scan(&id, &nama, &pengalaman)

		// Update manajer jika kandidat memiliki pengalaman lebih lama
		switch kota {
		case "1":
			if pengalaman > bandung.PengalamanKerja {
				bandung.ID = id
				bandung.Nama = nama
				bandung.PengalamanKerja = pengalaman
			}
		case "2":
			if pengalaman > jakarta.PengalamanKerja {
				jakarta.ID = id
				jakarta.Nama = nama
				jakarta.PengalamanKerja = pengalaman
			}
		case "3":
			if pengalaman > bekasi.PengalamanKerja {
				bekasi.ID = id
				bekasi.Nama = nama
				bekasi.PengalamanKerja = pengalaman
			}
		}
	}

	// Output data manajer untuk setiap kota
	fmt.Println(bandung.ID, bandung.Nama, bandung.PengalamanKerja)
	fmt.Println(jakarta.ID, jakarta.Nama, jakarta.PengalamanKerja)
	fmt.Println(bekasi.ID, bekasi.Nama, bekasi.PengalamanKerja)
}
