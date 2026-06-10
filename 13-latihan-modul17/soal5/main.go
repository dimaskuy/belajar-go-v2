package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var n int

	fmt.Print("Banyaknya topping: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Banyaknya topping harus lebih dari 0!")
		return
	}

	var topping int

	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-0.5)*(x-0.5)+(y-0.5)*(y-0.5) <= 0.5*0.5 {
			topping++
		}
	}

	fmt.Printf("Topping pada pizza: %d\n", topping)
	fmt.Printf("PI: %.10f\n", float64(topping)/float64(n)*4)
}
