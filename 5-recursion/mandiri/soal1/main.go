package main

import "fmt"

// fungsi rekursif
func hitungDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + hitungDigit(n/10)
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")

	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	fmt.Printf("Jumlah digit dari %d adalah %d\n", n, hitungDigit(n))
}