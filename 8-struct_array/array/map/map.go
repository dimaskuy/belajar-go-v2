package main

import "fmt"

/*
	map -> key: value
*/

func main() {
	var m1 = map[string]int{
		"satu": 1,
		"dua": 2,
		"tiga": 3,
	}
	var data = map[string]int{
		"Budi": 3123525,
		"Dimas": 12,
		"Jonathan": 12,
	}

	fmt.Println("Map 1:", m1)
	fmt.Println("Map 2:", data)

	fmt.Println("Data Seluruh Orang: ", data["Budi"], data["Dimas"], data["Jonathan"])
}