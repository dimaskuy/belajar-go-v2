package main

import "fmt"

func BiayaPOS() {
	var berat int
	var biayaSisa int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	kg := berat / 1000
	sisa := berat % 1000

	biayaKG := kg * 10000

	if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	totalBiaya := biayaKG

	if kg <= 10 {
		totalBiaya += biayaSisa
	}

	fmt.Println("Detail berat:", kg, "kg +", sisa, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKG, "+ Rp.", biayaSisa)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}

func BilanganFaktor() {
	var n int
	isValid := true
	isPrime := false
	faktor := 0
	fmt.Print("Bilangan: ")
	fmt.Scan(&n)

	if n <= 1 {
		isValid = false
	}

	if isValid {
		fmt.Println("Faktor: ")
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				fmt.Print(i, " ")
				faktor++
			}
		}
	} else {
		fmt.Println("Bilangan input tidak valid!")
	}

	if faktor == 2 {
		isPrime = true
	}

	fmt.Println("\nBilangan prima?", isPrime)
}

func main() {
	// BiayaPOS()
	BilanganFaktor()
}
