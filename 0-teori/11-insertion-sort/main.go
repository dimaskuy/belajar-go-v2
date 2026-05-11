package main

import "fmt"

func insertionAscending(data []int) []int {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j] > temp {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = temp
	}
	return data
}

func insertionDescending(data []int) []int {
	for i := 1; i < len(data); i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j] < temp {
			data[j+1] = data[j]
			j = j - 1
		}
		data[j+1] = temp
	}
	return data
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", arr)
	fmt.Println("Sorted array (Ascending):", insertionAscending(arr))
	fmt.Println("Sorted array (Descending):", insertionDescending(arr))
}