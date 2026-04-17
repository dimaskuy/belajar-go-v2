package main

import "fmt"

func hitungPoin(level int, i int) int {
	if level == 0 {
		return 0
	}
	fmt.Print("Level ke-", i, ": ")
	var poin int
	fmt.Scan(&poin)
	return poin + hitungPoin(level-1, i+1)
}

func main() {
	totalPoin := hitungPoin(5, 1)
	
	fmt.Println("Total poin: ", totalPoin)

	if totalPoin >= 500 {
		fmt.Println("Bonus: 50 poin")
		fmt.Println("Total akhir: ", totalPoin + 50)
	} else {
		fmt.Println("Bonus: 0 poin")
		fmt.Println("Total akhir: ", totalPoin)
	}
}