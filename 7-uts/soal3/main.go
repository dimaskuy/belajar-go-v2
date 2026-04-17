package main

import "fmt"

func hitungBiaya(berat int, total *int) {
	var beratAkhir int
	
	if berat <= 5 {
		beratAkhir = 10000
	} else {
		beratAkhir = 20000
	}
	
	*total += beratAkhir
}

func changeStatus(confirm string) bool {
	if confirm == "Y" || confirm == "y" {
		return true
	} else {
		return false
	}
}

func main() {
	var n, berat int
	var confirm string
	total := 0

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&berat)
		hitungBiaya(berat, &total)
	}

	fmt.Scan(&confirm)
	status := changeStatus(confirm)
	
	if status {
		s:="TERKIRIM"
		fmt.Printf("Total Biaya: %d\nStatus: %s", total, s)
	} else {
		fmt.Println("Pengiriman dibatalkan")
	}
	
}