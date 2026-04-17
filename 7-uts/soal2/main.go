package main

import "fmt"

func tarifKendaraan(kendaraan string) int {
	switch kendaraan {
	case "mobil":
		return 5000
	case "motor":
		return 2000
	}
	return 0
}

func totalBiaya(tarif int, durasi int, akhir *int) {
	*akhir += tarif * durasi
}

func main() {
	var n int
	var kendaraan string
	var durasi int
	akhir := 0
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&kendaraan, &durasi) // format "[kendaraan] [durasi]"
		tarif:=tarifKendaraan(kendaraan)
		totalBiaya(tarif, durasi, &akhir)
	}
	fmt.Println("Total biaya parkir: ", akhir)
}