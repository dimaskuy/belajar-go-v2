package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	// (f o g o h) (a)
	fmt.Println(f(g(h(a))))
	// (g o h o f) (b)
	fmt.Println(g(h(f(b))))
	// (h o f o g) (c)
	fmt.Println(h(f(g(c))))
}
