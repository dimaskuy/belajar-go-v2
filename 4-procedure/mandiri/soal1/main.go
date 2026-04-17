package main

import "fmt"

// prosedur hitungDiskon
func hitungDiskon(totalBelanja int) {
	if totalBelanja >= 500000 {
		fmt.Println("Diskon: 20%")
		fmt.Println("Total bayar: ", totalBelanja - (totalBelanja * 20 / 100))
	} else if totalBelanja >= 300000 {
		fmt.Println("Diskon: 15%")
		fmt.Println("Total bayar: ", totalBelanja - (totalBelanja * 15 / 100))
	} else if totalBelanja >= 100000 {
		fmt.Println("Diskon: 10%")
		fmt.Println("Total bayar: ", totalBelanja - (totalBelanja * 10 / 100))
	} else {
		fmt.Println("Diskon: Tidak ada diskon")
	}
}

func main() {
	var nama string
	var totalBelanja int
	fmt.Print("Nama pelanggan: ")
	fmt.Scan(&nama)
	fmt.Print("Total belanja: ")
	fmt.Scan(&totalBelanja)
	
	fmt.Println("Pelanggan: ", nama)
	hitungDiskon(totalBelanja)
}