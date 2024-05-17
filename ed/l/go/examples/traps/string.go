package main

import (
	"fmt"
)

func main() {
	lenExample()
	loopExample()
}

func lenExample() {
	s := "汉"
	fmt.Printf("stringLe: =%d \n", len(s)) // 3

	r := []rune(s)
	fmt.Printf("runeLen: %d \n", len(r)) // 1
}

func loopExample() {
	fmt.Printf("\nloopAsString: \n")
	s := "hêllo"
	for i, r := range s {
		fmt.Printf("position %d: %c\n", i, r)
	}
	// Result:
	// position 0: h
	// position 1: ê
	// position 3: l
	// position 4: l
	// position 5: o

	fmt.Printf("\nloopAsRuneSlice: \n")
	rs := []rune(s)
	for i, r := range rs {
		fmt.Printf("position %d: %c\n", i, r)
	}
	// Result:
	// position 0: h
	// position 1: ê
	// position 2: l
	// position 3: l
	// position 4: o
}
