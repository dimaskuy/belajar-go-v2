package main

import (
	"fmt"
)

func menghitungNilaiAkhir(t, uts, uas int) float64 {
	return (float64(t)*0.3)+(float64(uts)*0.3)+(float64(uas)*0.4)
}

func tentukanGrade(akhir float64) {
	var grade string
	if akhir >= 85 {
		grade = "A"
	} else if akhir >= 70 {
		grade = "B"
	} else if akhir >= 55 {
		grade = "C"
	} else if akhir >= 40 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Printf("Nilai Akhir: %.2f\nGrade: %s", akhir, grade)
}

func main() {
	var t, uts, uas int

	fmt.Scanln(&t)
	fmt.Scanln(&uts)
	fmt.Scanln(&uas)
	akhir:=menghitungNilaiAkhir(t,uts,uas)
	tentukanGrade(akhir)
}


