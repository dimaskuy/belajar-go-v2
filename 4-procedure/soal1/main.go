package main

import "fmt"

func bebekDengklek(n int) {
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			if j == i || j == 2*n-2-i {
				fmt.Print("*")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	bebekDengklek(2)
	bebekDengklek(10)
}
