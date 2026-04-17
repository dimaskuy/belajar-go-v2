// Dimas IF-05-02
package main

import "fmt"

func faktorial(n int) int {
	if n == 0 {
		return 1
	} 
		return n * faktorial(n-1)
}

func ganjilGenap(n int) string {
	if n%2 == 0 {
		return "genap"
	}
		return "ganjil"
}

// PROSEDUR
func formatFaktorial(hasil int, output *string) {
	*output = fmt.Sprintf("Faktorial: %d", hasil)
}

func formatGG(jenis string, output *string) {
	*output = fmt.Sprintf("Jenis: %s", jenis)
}

func main() {
	var n int
	var hasilFaktorial, hasilGG string

	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	faktorial := faktorial(n)
	jenis := ganjilGenap(n)

	formatFaktorial(faktorial, &hasilFaktorial)
	formatGG(jenis, &hasilGG)

	fmt.Println(hasilFaktorial)
	fmt.Println(hasilGG)
}
