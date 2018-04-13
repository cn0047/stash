package main

import "fmt"

func main() {
	a := 5.0
	b := int(a)
	c := uint(b)
	fmt.Printf("\n a = %T, b = %T, c = %T", a, b, c)
}
