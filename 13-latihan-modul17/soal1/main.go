package main

import "fmt"

func main() {
	const sen = 9999.0

	var bil float64
	var total float64
	var jumlahData int

	for {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&bil)

		if bil == sen {
			break
		}

		total += bil
		jumlahData++
	}

	if jumlahData == 0 {
		fmt.Println("Tidak ada data yang dimasukkan")
		return
	}

	rerata := total / float64(jumlahData)
	fmt.Printf("Rerata: %.2f\n", rerata)
}
