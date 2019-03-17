package main

import (
	"fmt"
	"geometry"
)

func main() {
	var length, breadth = 10, 20
	fmt.Println("Area is", geometry.Area(length, breadth))
	fmt.Println("Perimeter is", geometry.Perimeter(length, breadth))
}
