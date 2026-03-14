package main

import "fmt"

func nilaiAkhir(t, uts, uas int) int {
	return int(float64(t)*0.3 + float64(uts)*0.3 + float64(uas)*0.4 + 0.5)
}

func nilaiHuruf(n int) string {
	var grade string
	if n >= 80 && n <= 100 {
		grade = "A"
	} else if n >= 70 {
		grade = "B"
	} else if n >= 60 {
		grade = "C"
	} else if n >= 50 {
		grade = "D"
	} else if n <= 50 && n >= 0 {
		grade = "E"
	} else {
		grade = "ERR."
	}

	return grade
}

func main() {
	var nama string
	var t, uts, uas int
	fmt.Print("Nama: ")
	fmt.Scan(&nama)

	fmt.Print("Nilai tugas: ")
	fmt.Scan(&t)
	fmt.Print("Nilai UTS: ")
	fmt.Scan(&uts)
	fmt.Print("Nilai UAS: ")
	fmt.Scan(&uas)

	akhir := nilaiAkhir(t, uts, uas)
	huruf := nilaiHuruf(akhir)

	fmt.Println("Nama Mahasiswa: ", nama)
	fmt.Println("Nilai Akhir: ", akhir)
	fmt.Println("Nilai Huruf: ", huruf)

}
