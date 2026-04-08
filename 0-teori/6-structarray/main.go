package main

import "fmt"

type geometry struct {
	area      int
	perimeter int
}

type rectangle struct {
	length   int
	width    int
	color    string
	property geometry
}

func input(obj *rectangle) {
	fmt.Print("Length: ")
	fmt.Scan(&obj.length)
	fmt.Print("Width: ")
	fmt.Scan(&obj.width)
	fmt.Print("Color: ")
	fmt.Scan(&obj.color)
}

func main() {
	var persegi rectangle
	// persegi.length = 4
	// persegi.width = 2
	// persegi.color = "blue"
	input(&persegi)
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = 2 * (persegi.length + persegi.width)

	fmt.Println("Luas: ", persegi.property.area)
	fmt.Println("Parimeter: ", persegi.property.perimeter)
	fmt.Println("Color: ", persegi.color)
}
