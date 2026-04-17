package main

import "fmt"

func hitungLuas(pilihan int) {
	switch pilihan {
	case 1:
		var sisi int
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&sisi)
		fmt.Println("Luas bangun datar persegi: ", sisi * sisi)
	case 2:
		var panjang, lebar int
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		fmt.Println("Luas bangun datar persegi panjang: ", panjang * lebar)
	case 3:
		var alas, tinggi int
		fmt.Print("Masukkan alas: ")
		fmt.Scan(&alas)
		fmt.Print("Masukkan tinggi: ")
		fmt.Scan(&tinggi)
		fmt.Println("Luas bangun datar segitiga: ", 0.5 * float64(alas) * float64(tinggi))
	default:
		fmt.Println("Pilihan tidak ada")
	}
}

func main() {
	var pilihan int
	fmt.Print("Pilihan bangun datar (1/2/3): ")
	fmt.Scan(&pilihan)
	hitungLuas(pilihan)
}