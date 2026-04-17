// Dimas IF-05-02
package main

import "fmt"

func biayaParkir(jam int) int {
	return jam * 3000
}

func hitungDenda(jam int) int {
	if jam > 5 {
		return 2000
	}
	return 0
}

func totalAkhir(biaya int, denda int) int {
	return biaya + denda
}

func statusParkir(jam int) string {
	if jam > 5 {
		return "TERLAMBAT"
	}
	return "TEPAT WAKTU"
}

func main() {
	// input & output
	var jam int
	var totalBiaya, denda, akhir int
	var status string

	fmt.Print("Masukkan jumlah jam parkir: ")
	fmt.Scanln(&jam)

	totalBiaya = biayaParkir(jam)
	denda = hitungDenda(jam)
	akhir = totalAkhir(totalBiaya, denda)
	status = statusParkir(jam)

	fmt.Println("Biaya: ", totalBiaya)
	fmt.Println("Denda: ", denda)
	fmt.Println("Total: ", akhir)
	fmt.Println("Status: ", status)
}