package main

import "fmt"

func main() {
	for {
		var massa, volume, mr float64
		fmt.Scan(&massa, &volume, &mr)

		if massa == 0 {
			break
		}

		if massa < 0 || volume <= 0 || mr <= 0 {
			fmt.Println(0)
		} else {
			mol := massa / mr
			L := volume / 1000
			M := mol / L
			fmt.Printf("%.2f\n", M)
		}
	}
}
