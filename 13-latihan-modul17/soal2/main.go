package main

import "fmt"

func main() {
	var x string
	var n int

	fmt.Print("Masukkan string x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah data n: ")
	fmt.Scan(&n)

	data := make([]string, 0, n)

	for i := 0; i < n; i++ {
		var dataString string
		fmt.Print("Masukkan data ", i+1, ": ")
		fmt.Scan(&dataString)
		data = append(data, dataString)
	}

	var ditemukan bool
	var posisi int
	var jumlahData int

	for i, item := range data {
		if item == x {
			jumlahData++
			if !ditemukan {
				ditemukan = true
				posisi = i + 1
			}
		}
	}

	adaDua := jumlahData >= 2

	fmt.Println("a. Apakah string x ada dalam kumpulan n data string tersebut?", ditemukan)
	fmt.Println("b. Pada posisi ke berapa string x tersebut ditemukan?", posisi)
	fmt.Println("c. Ada berapakah string x dalam kumpulan n data string tersebut?", jumlahData)
	fmt.Println("d. Adakah sedikitnya dua string x dalam n data string tersebut?", adaDua)
}
