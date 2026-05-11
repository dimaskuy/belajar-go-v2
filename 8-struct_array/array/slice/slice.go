package main

import "fmt"

/*
	slice:
	- pointer -> array
	- length
	- capacity

	slicing:
	misal, arr[start:end-1]
	misal, arr[:] -> ambil semua
	misal, arr[:end-1] -> ambil dari awal sampai end-1
	misal, arr[start:] -> ambil dari start sampai akhir
*/

func main() {
	var myArr = [3]int{1, 2, 3} // Array (statis, kita menentukan index-nya)
	var mySlice = []int{1, 2, 3, 4, 5} // Slice (dinamis)

	fmt.Println("Array:", myArr)
	fmt.Println("Slice:", mySlice)

	mySlice = append(mySlice, 6, 7, 8, 9, 10)
	fmt.Println("Slice:", mySlice)
}