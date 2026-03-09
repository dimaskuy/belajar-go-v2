package main

import "fmt"

/* PASS BY VALUE & PASS BY REFERENCE */

// * = mengakses nilai dari alamat yang akan dikirim nanti
//

func pecahUang(uang int, k10 *int, k2 *int, k1 *int) {
	*k10 = uang / 10000
	uang = uang % 10000

	*k2 = uang / 2000
	uang = uang % 2000

	*k1 = uang / 1000
}

func konversiSuhu(c float64, r *float64, f *float64, k *float64) {
	*r = (4.0 / 5.0) * c
	*f = (9.0/5.0)*c + 32.0
	*k = c + 273.15
}

func main() {
	// konversi uang
	uang := 27000
	var k10, k2, k1 int
	pecahUang(uang, &k10, &k2, &k1)

	fmt.Println("10000 : ", k10)
	fmt.Println("2000 : ", k2)
	fmt.Println("1000 : ", k1)

	// konversi suhu
	var c float64 = 25
	var r, f, k float64

	konversiSuhu(c, &r, &f, &k)

	fmt.Println("Reamur: ", r)
	fmt.Println("Farenheit: ", f)
	fmt.Println("Kelvin: ", k)
}
