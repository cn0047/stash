package main

import (
	"fmt"
)

func main() {
	f1()
}

func g1[myType any](x myType) {
	fmt.Printf("[g1] x value:%#v; x type: %T \n", x, x) // [g1] x value:1; x type: int
}

func f1() {
	x := 1
	g1(x)
}
