// Dimas IF-05-02
package main

import "fmt"

// Fungsi 1: Menghitung total (liter x harga).
func hitungTotal(liter int, harga int) int {
	return liter * harga
}

// Fungsi 2: Menghitung cashback (jika total > 50000, cashback = 5% x total).
func hitungCashback(total int) int {
	if total > 50000 {
		return total * 5 / 100
	}
	return 0
}

// Prosedur 1: Menghitung bayar akhir (Total - Cashback).
func hitungBayarAkhir(total int, cashback int) int {
	return total - cashback
}

// Prosedur 2: Inisialisasi data transaksi.
func printData(liter int, harga int, total int, cashback int, bayarAkhir int) {
	fmt.Println("Total: ", total)
	fmt.Println("Cashback: ", cashback)
	fmt.Println("Total Bayar: ", bayarAkhir)
}

func main() {
	var liter, harga, total, cashback, bayarAkhir int
	fmt.Print("Masukkan liter: ")
	fmt.Scanln(&liter)
	fmt.Print("Masukkan harga: ")
	fmt.Scanln(&harga)
	total = hitungTotal(liter, harga)
	cashback = hitungCashback(total)
	bayarAkhir = hitungBayarAkhir(total, cashback)
	printData(liter, harga, total, cashback, bayarAkhir)
}