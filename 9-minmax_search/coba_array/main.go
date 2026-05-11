package main

import "fmt"

type myArr []int

func smaller(arr myArr, target int) int {
	min := arr[0]
	var i int

	for i < target {
		if min > arr[i] {
			min = arr[i]
		}
		i++
	}
	return min
}

func main() {
	arr := myArr{2, 3, 4, 10, 40}
	fmt.Println(smaller(arr, 4))
}