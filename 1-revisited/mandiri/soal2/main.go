package main

import "fmt"

func main() {
	var snack, kapasitas int
	fmt.Scan(&snack)

	if snack == 0 {
		fmt.Println()
		return
	}

	for n := 1; n <= 5; n++ {
		if snack <= 0 {
			break
		}

		kapasitas = 5 + (n * 5)

		if snack >= kapasitas {
			fmt.Print(kapasitas, " ")
			snack -= kapasitas
		} else {
			fmt.Print(snack, " ")
			snack = 0
		}
	}
}
