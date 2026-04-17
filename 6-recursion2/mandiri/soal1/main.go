package main

import "fmt"

// prosedur untuk mengecek harga barang
func cekHarga(pilihan int) int {
	switch pilihan {
	case 1:
		return 5000
	case 2:
		return 7000
	case 3:
		return 2500
	case 4:
		return 4000
	case 5:
		return 2000
	default:
		return 0
	}
}
// prosedur untuk menampilkan menu
func tampilkanMenu() {
	fmt.Println("Daftar Barang:")
	fmt.Println("1. Susu kotak = Rp 5.000")
	fmt.Println("2. Coklat batang = Rp 7.000")
	fmt.Println("3. Mie instan = Rp 2.500")
	fmt.Println("4. Teh botol = Rp 4.000")
	fmt.Println("5. Kardus bekas = Rp 2.000")
	fmt.Println("0. Berhenti")
}

func main() {
	totalHarga := 0

	tampilkanMenu()
	for i := 1; i <= 3; i++ {
		fmt.Print("Masukkan barang ke-", i, ": ")
		var pilihan int
		fmt.Scanln(&pilihan)
		if pilihan == 0 {
			break
		}
		totalHarga += cekHarga(pilihan)
	}
	
	fmt.Println("Total harga: Rp", totalHarga)
}