package main

import "fmt"

// insertion sort
func main() {
	var arr []int
	var n int
	i := 1
	jarak := 0

	for {
		var input int
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		arr = append(arr, input)
	}

	n = len(arr)
	for i <= n-1 {
		j := i
		temp := arr[j]
		for j > 0 && temp < arr[j-1] {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = temp
		jarak = arr[i] - arr[i-1]
		i++
	}

	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	if jarak != arr[1]-arr[0] {
		fmt.Print("Data berjarak tidak tetap\n")
	} else {
		fmt.Printf("Data berjarak %d\n", jarak)
	}
}
