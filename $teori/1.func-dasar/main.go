package main

import "fmt"

func konversiWaktu(jam, menit, detik int) int {
	convJam := jam * 3600
	convMnt := menit * 60
	return convJam + convMnt + detik
}

func cekSegitiga(d1, d2, d3 int) string {
	valid := d1+d2 > d3 && d1+d3 > d2 && d2+d3 > d1
	if valid {
		return "segitiga"
	}
	return "bukan segitiga"
}

func itungDiskon(prc int, isMember bool) int {
	if prc > 100000 && isMember {
		return prc - int(float64(prc)*0.10)
	} else if prc > 100000 && !isMember {
		return prc - int(float64(prc)*0.05)
	}
	return prc
}

func fibo(n int) int {
	a := 0
	b := 1
	for i := 1; i <= n; i++ {
		next := a + b
		a = b
		b = next
	}
	return b
}

func main() {
	fmt.Println(konversiWaktu(0, 3, 10))
	fmt.Println(cekSegitiga(1, 10, 5))
	fmt.Println(itungDiskon(300000, true))
	fmt.Println(fibo(7))
}
