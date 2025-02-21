package main

import (
	"fmt"
)

func main() {
	simpleSet()
}

func simpleSet() {
	type Set[T comparable] = map[T]bool
	set := Set[string]{"one": true, "two": false}

	fmt.Printf("one: %v\n", set["one"])
	fmt.Printf("six: %v\n", set["six"])
	fmt.Printf("set: %T\n", set) // set: map[string]bool
}
