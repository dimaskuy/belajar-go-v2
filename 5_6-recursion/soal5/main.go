// Dimas IF-05-02
package main

import "fmt"

func hitungDigit(n int) int {
	if n < 10 {
		return 1
	}
	return 1 + hitungDigit(n/10)
}

func balikRekursif(n int, hasil int) int {
	if n == 0 {
		return hasil
	}
	return balikRekursif(n/10, hasil*10+(n%10))
}

func balikAngka(n int) int {
	return balikRekursif(n, 0)
}

func formatDigit(digit int, out *int) {
	*out = digit
}

func formatBalik(balik int, out *int) {
	*out = balik
}

func main() {
	var n int
	var hasilDigit, hasilBalik int

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	d := hitungDigit(n)
	b := balikAngka(n)

	formatDigit(d, &hasilDigit)
	formatBalik(b, &hasilBalik)

	fmt.Println("Jumlah digit: ", hasilDigit)
	fmt.Println("Reverse: ", hasilBalik)
}