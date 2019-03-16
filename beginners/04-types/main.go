package main

import (
	"fmt"
)

func main() {
	var i int16 = 55     //int16
	var j float32 = 67.8 //float32
	sum := i + int16(j)  //j is converted to int
	fmt.Println(sum)

	const aa int = 48
	const bb bool = true
	fmt.Println(aa, bb)
}
