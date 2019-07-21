package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	s := a[1:3]
	fmt.Printf("array: %+v; slice: %+v \n", a, s) // array: [1 2 3]; slice: [2 3]

	// change array
	a[1] = 9
	fmt.Printf("array: %+v; slice: %+v \n", a, s) // array: [1 9 3]; slice: [9 3]

	// change slice
	s[1] = 7
	fmt.Printf("array: %+v; slice: %+v \n", a, s) // array: [1 9 7]; slice: [9 7]
}
