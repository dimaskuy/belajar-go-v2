package main

import "fmt"

func team(namaA, namaB *string) {
	fmt.Print("Tim A: ")
	fmt.Scan(namaA)
	fmt.Print("Tim B: ")
	fmt.Scan(namaB)
}

func score(namaA, namaB string, winners *[]string) {
	var scoreA, scoreB int

	for {
		fmt.Printf("Pertandingan %d: ", len(*winners)+1)
		fmt.Scan(&scoreA, &scoreB)

		if scoreA < 0 || scoreB < 0 {
			break
		}

		if scoreA > scoreB {
			*winners = append(*winners, namaA)
		} else if scoreB > scoreA {
			*winners = append(*winners, namaB)
		}
	}
}

func winner(winners []string) {
	for i := 0; i < len(winners); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, winners[i])
	}
}

func main() {
	var namaA, namaB string
	var winners []string

	team(&namaA, &namaB)
	score(namaA, namaB, &winners)
	winner(winners)
	fmt.Println("Pertandingan Selesai!")
}
