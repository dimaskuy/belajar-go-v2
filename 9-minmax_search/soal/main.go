package main

import "fmt"

// TODO: user  memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	/* I.S. Terdefinisi array dinamis arrBerat
	Proses: Menghitung berat minimum dan maksimum dalam array
	F.S. Menampilkan berat minimum dan maksimum balita */
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	sum := 0.0
	for i := 0; i < 100; i++ {
		sum += arrBerat[i]
	}
	return sum / float64(n)
}

func main() {
	var balita arrBalita
	var minimum float64
	var maximum float64
	var rata float64
	var n int

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Masukan berat balita ke-", i+1, ": ")
		fmt.Scan(&balita[i])
	}

	hitungMinMax(balita, n, &minimum, &maximum)
	rata = rerata(balita, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", minimum)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", maximum)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata)
}
