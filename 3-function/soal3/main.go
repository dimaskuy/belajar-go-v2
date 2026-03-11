package main

import "fmt"

func isInsideCircle(x, y, xc, yc, r int) bool {
	// rumus jarak dari titik ke pusat lingkaran
	return ((x-xc)*(x-xc))+((y-yc)*(y-yc)) <= (r * r)
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var x, y int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	inside1 := isInsideCircle(x, y, x1, y1, r1)
	inside2 := isInsideCircle(x, y, x2, y2, r2)

	if inside1 && inside2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inside1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inside2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
