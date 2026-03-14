package main

import "fmt"

func biayaParkir(durasi int) int {
	biaya := 5000
	if durasi > 1 {
		biaya += (durasi - 1) * 3000
	} else if durasi < 0 {
		return 0
	}
	return biaya
}

func checkDiskon(durasi, biaya int) int {
	if durasi > 5 {
		return int(float64(biaya) * 0.9) // biaya - (biaya*0.1)
	}
	return biaya
}

func main() {
	var nama string
	var durasi int

	fmt.Print("Kendaraan: ")
	fmt.Scanln(&nama)
	fmt.Print("Lama Parkir: ")
	fmt.Scanln(&durasi)

	biaya := checkDiskon(durasi, biayaParkir(durasi))
	fmt.Println("Kendaraan", nama)
	fmt.Println("Lama Parkir: ", durasi)
	fmt.Println("Total Biaya Parkir: ", biaya)
}
