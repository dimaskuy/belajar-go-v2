// Dimas IF-05-02
package main

import "fmt"

// Fungsi 1: Menghitung biaya dasar (berat x 10000) dan ditambah biaya layanan jika Express.
func biayaDasar(berat int, jenis string) int {
	biaya := (berat * 10000)
	if jenis == "express" {
		return biaya + 5000
	} else {
		return biaya
	}
}

// Fungsi 2: Menentukan kategori string ("Murah" atau "Mahal") berdasarkan total biaya
func kategori(biaya int) string {
	if biaya > 30000 {
		return "mahal"
	} else {
		return "murah"
	}
}

//  Prosedur 1: Menghitung nilai diskon (cek jika biaya > 50000).
func cekDiskon(biaya int, isDiskon *bool) {
	*isDiskon = biaya > 50000
}

// Prosedur 2: Menghitung total akhir setelah diskon.
func diskon(biaya int, jumlahDiskon *int, terdiskon *int) {
	var isDiskon bool
	cekDiskon(biaya, &isDiskon)
	if isDiskon {
		*terdiskon = biaya - 5000
		*jumlahDiskon = 5000
	} else {
		*terdiskon = biaya
		*jumlahDiskon = 0
	}
}

func main() {
	var berat, jumlahDiskon, terdiskon int
	var jenis string

	fmt.Print("Berat (kg): ")
	fmt.Scanln(&berat)
	fmt.Print("Jenis layanan (reguler/express): ")
	fmt.Scanln(&jenis)

	biaya := biayaDasar(berat, jenis)
	kategori := kategori(biaya)
	diskon(biaya, &jumlahDiskon, &terdiskon)

	fmt.Println("Total: Rp", biaya)
	fmt.Println("Kategori: ", kategori)
	fmt.Println("Diskon: Rp", jumlahDiskon)
	fmt.Println("Total Bayar: Rp", terdiskon)
}
