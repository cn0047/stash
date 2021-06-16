package main

import (
	"fmt"
)

func main() {
	varAfterVar()
}

func varAfterVar() {
	var (
		a = "foo"
		b = a
	)
	fmt.Printf("b = %s\n", b) // b = foo
}
