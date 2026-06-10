package main

import "fmt"

func main() {
	var n int

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	sum := 0.0

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			sum -= 1.0 / float64(2*i-1)
		} else {
			sum += 1.0 / float64(2*i-1)
		}
	}

	sum = 0.0

	for i := 1; ; i++ {
		sebelum := sum

		if i%2 == 0 {
			sum -= 1.0 / float64(2*i-1)
		} else {
			sum += 1.0 / float64(2*i-1)
		}

		selisih := 4.0 / float64(2*i-1)

		if selisih < 0.00001 {
			fmt.Printf("Hasil PI: %.10f\n", 4*sebelum)
			fmt.Printf("Hasil PI: %.10f\n", 4*sum)
			fmt.Printf("Pada i ke: %d\n", i)
			break
		}
	}
}
