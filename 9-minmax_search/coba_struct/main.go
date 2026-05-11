package main

import "fmt"

type dataMhs struct {
    nama string
    npm  string
    ipk  float64
}

type arrMhs []dataMhs

func IPKIndex(T arrMhs, n int) *dataMhs {
    idx := 0
    for j := 1; j < n; j++ {
        if T[idx].ipk < T[j].ipk {
            idx = j
        }
    }
    return &T[idx]
}

func main() {
    var data arrMhs
    var n int

    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Scan(&data[i].nama)
        fmt.Scan(&data[i].npm)
        fmt.Scan(&data[i].ipk)
    }

    mhs := IPKIndex(data, n)
    
    fmt.Println(mhs.nama)
    fmt.Println(mhs.npm)
    fmt.Println(mhs.ipk)
}