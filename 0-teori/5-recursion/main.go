package main

import "fmt"

// FPB
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a < 0 || b < 0 {
		return
	}
	hasil := gcd(a, b)
	fmt.Println(hasil)
}
