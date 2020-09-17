package main

import (
	"fmt"
)

func main() {
	one()
}

func one() {
	a := 5.0
	b := int(a)
	c := uint(b)
	fmt.Printf("\n a = %T, b = %T, c = %T", a, b, c) //  a = float64, b = int, c = uint
}
