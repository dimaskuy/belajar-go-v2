// Dimas IF-05-02
package main

import "fmt"

func totalBelanja(b1, b2, b3, h1, h2, h3 int) int {
	return (b1 * h1) + (b2 * h2) + (b3 * h3)
}

func diskon(total int) int {
	if total > 100000 {
		return total * 10 / 100
	} else {
		return 0
	}
}

func totalAkhir(total int, diskon int, akhir *int) {
	*akhir = total - diskon
}

func validasiInput(b1, b2, b3, h1, h2, h3 int, valid *bool) {
	*valid = b1 > 0 && b2 > 0 && b3 > 0 && h1 > 0 && h2 > 0 && h3 > 0
}

func main() {
	var b1, b2, b3, h1, h2, h3 int
	var isValid bool
	var akhir int
	fmt.Print("Barang 1: ")
	fmt.Scanln(&b1, &h1)
	fmt.Print("Barang 2: ")
	fmt.Scanln(&b2, &h2)
	fmt.Print("Barang 3: ")
	fmt.Scanln(&b3, &h3)

	validasiInput(b1, b2, b3, h1, h2, h3, &isValid)

	if isValid {
		total := totalBelanja(b1, b2, b3, h1, h2, h3)
		diskon := diskon(total)
		totalAkhir(total, diskon, &akhir)
		fmt.Println("Total Belanja: ", total)
		fmt.Println("Diskon: ", diskon)
		fmt.Println("Total Akhir: ", akhir)
	} else {
		fmt.Println("Input tidak valid")
	}

}
