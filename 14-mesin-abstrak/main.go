// MESIN ABSTRAK/DOMINO || SOAL LATIHAN MODUL 18
package main

import (
	"fmt"
	"math/rand/v2"
)

/*
		IMPLEMENTASI
	 1. Implementasi operasi dasar mesin domino sebagai sebuah subprogram:
	    a) Buat tipe data kartu domino (Domino) yang menyimpan informasi
	    Ø gambar (suit) kedua sisi kartu
	    Ø nilai kartu
	    Ø Boolean data yang menyatakan kartu ini balak atau bukan
	    Ø Buat tipe data satu set kartu domino (Dominoes)
	    Ø Array menyimpan 28 kartu Domino
	    Ø Jumlah kartu tersisa dalam array tersebut
	    b) prosedur kocokKartu(Dominoes)
	    c) fungsi ambilKartu(Dominoes) à Domino
	    d) fungsi gambarKartu(Domino,suit int) à int
	    e) fungsi nilaiKartu(Domino) à int

	 2. Realisasi aksi berikut menggunakan operasi-operasi dasar mesin domino:
	    a) prosedur galiKartu(Dominoes,Domino) yang mengambil kartu dari tumpukan sampai
	    diperoleh kartu dengan gambar (suit) yang sama dengan kartu yang diberikan
	    b) fungsi sepasangKartu(Domino,Domino) à boolean; yang memberikan nilai true jika total nilai kartu adalah 12 dan false jika tidak.
	 3. Implementasi salah satu permainan domino. Lihat lampiran untuk deskripsi permainan Gapleh.
	 4. Implementasi mesin abstrak karakter yang bekerja terhadap untaian karakter (yang diakhiri dengan penanda titik (".") dan mempunyai sejumlah operasi dasar.
	    a) Operasi dasar mesin karakter:
	    Ø Prosedur start(); yang menyiapkan mesin karakter di awal rangkaian karakter.
	    Ø Prosedur maju(); yang memajukan pembaca ke posisi karakter berikutnya.
	    Ø Fungsi eop(); yang mengembalikan nilai true apabila sudah mencapai akhir rangkaian, sampai ke penanda titik (".").
	    Ø Fungsi cc(); yang mengembalikan karakter yang sedang terbaca, atau berada pada posisi pembacaan mesin.
	    b) Dengan operasi dasar di atas buat algoritma untuk:
	    Ø Membaca seluruh karakter yang diberikan ke mesin karakter tersebut.
	    Ø Menghitung berapa banyak karakter yang terbaca.
	    Ø Menghitung ada berapa huruf "A" yang terbaca.
	    Ø Menghitung frekuensi kemunculan huruf "A" terhadap seluruh karakter terbaca.
	    Ø Menghitung ada berapa kata "LE" (pasangan berturutan huruf "L" dan "E") yang terbaca.
*/
type Domino struct {
	kiri  int
	kanan int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu [28]Domino
	n     int
}

func makeDomino(kiri, kanan int) Domino {
	return Domino{
		kiri:  kiri,
		kanan: kanan,
		nilai: kiri + kanan,
		balak: kiri == kanan,
	}
}

func initDominoes(D *Dominoes) {
	idx := 0

	for i := 0; i <= 6; i++ {
		for j := i; j <= 6; j++ {
			D.kartu[idx] = makeDomino(i, j)
			idx++
		}
	}

	D.n = 28
}

func kocokKartu(D *Dominoes) {
	for i := 0; i < D.n; i++ {
		j := rand.IntN(D.n)
		D.kartu[i], D.kartu[j] = D.kartu[j], D.kartu[i]
	}
}

func ambilKartu(D *Dominoes) Domino {
	if D.n == 0 {
		return Domino{}
	}

	D.n--
	return D.kartu[D.n]
}

func gambarKartu(d Domino, suit int) int {
	if d.kiri == suit {
		return d.kanan
	}

	if d.kanan == suit {
		return d.kiri
	}

	return -1
}

func nilaiKartu(d Domino) int {
	return d.nilai
}

func cetakDomino(d Domino) {
	fmt.Printf("(%d,%d)", d.kiri, d.kanan)
}

func samaSuit(a, b Domino) bool {
	return a.kiri == b.kiri ||
		a.kiri == b.kanan ||
		a.kanan == b.kiri ||
		a.kanan == b.kanan
}

func galiKartu(D *Dominoes, target Domino) Domino {
	fmt.Println("\n=== GALI KARTU ===")

	for D.n > 0 {
		kartu := ambilKartu(D)

		fmt.Print("Mengambil ")
		cetakDomino(kartu)

		if samaSuit(kartu, target) {
			fmt.Println(" -> COCOK")
			return kartu
		}

		fmt.Println(" -> TIDAK COCOK")
	}

	fmt.Println("Tidak ditemukan kartu yang cocok")
	return Domino{}
}

func sepasangKartu(a, b Domino) bool {
	return nilaiKartu(a)+nilaiKartu(b) == 12
}

// MESIN KARAKTER

var pita string
var idx int
var currChar byte

func start() {
	idx = 0

	if len(pita) > 0 {
		currChar = pita[idx]
	}
}

func maju() {
	idx++

	if idx < len(pita) {
		currChar = pita[idx]
	} else {
		currChar = '.'
	}
}

func eop() bool {
	return currChar == '.'
}

func cc() byte {
	return currChar
}

// SOAL

func bacaPita() {
	start()

	fmt.Print("Isi pita : ")

	for !eop() {
		fmt.Printf("%c", cc())
		maju()
	}

	fmt.Println()
}

func hitungKarakter() int {
	start()

	jumlah := 0

	for !eop() {
		jumlah++
		maju()
	}

	return jumlah
}

func hitungA() int {
	start()

	jumlah := 0

	for !eop() {
		if cc() == 'A' {
			jumlah++
		}
		maju()
	}

	return jumlah
}

func freqA() float64 {
	total := hitungKarakter()

	if total == 0 {
		return 0
	}

	return float64(hitungA()) / float64(total)
}

func hitungLE() int {
	start()

	jumlah := 0
	prev := byte(0)

	for !eop() {
		if prev == 'L' && cc() == 'E' {
			jumlah++
		}

		prev = cc()
		maju()
	}

	return jumlah
}

func main() {
	var deck Dominoes

	initDominoes(&deck)

	fmt.Println("~~ MESIN DOMINO ~~")
	fmt.Println("Jumlah kartu awal:", deck.n)

	kocokKartu(&deck)

	kartu1 := ambilKartu(&deck)
	kartu2 := ambilKartu(&deck)

	fmt.Print("Kartu 1: ")
	cetakDomino(kartu1)
	fmt.Println()

	fmt.Print("Kartu 2: ")
	cetakDomino(kartu2)
	fmt.Println()

	fmt.Println("Nilai kartu 1:", nilaiKartu(kartu1))
	fmt.Println("Nilai kartu 2:", nilaiKartu(kartu2))

	if kartu1.balak {
		fmt.Println("Kartu 1 adalah BALAK")
	}

	if kartu2.balak {
		fmt.Println("Kartu 2 adalah BALAK")
	}

	if sepasangKartu(kartu1, kartu2) {
		fmt.Println("Total nilai kedua kartu = 12")
	} else {
		fmt.Println("Total nilai kedua kartu != 12")
	}

	galiKartu(&deck, kartu1)

	fmt.Println("Sisa kartu:", deck.n)

	fmt.Println("\n~~ MESIN KARAKTER ~~")

	fmt.Print("Masukkan pita (akhiri dengan titik): ")
	fmt.Scanln(&pita)

	if len(pita) == 0 || pita[len(pita)-1] != '.' {
		pita += "."
	}

	bacaPita()

	fmt.Println("Jumlah karakter: ", hitungKarakter())
	fmt.Println("Jumlah huruf A: ", hitungA())
	fmt.Printf("Frekuensi A: %.2f\n", freqA())
	fmt.Println("Jumlah kata LE: ", hitungLE())
}
