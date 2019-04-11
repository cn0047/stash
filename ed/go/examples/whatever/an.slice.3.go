package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 2) // length
	s1 = append(s1, 1)
	fmt.Println(s1, len(s1), cap(s1)) // [0 0 1] 3 4

	s2 := make([]int, 0, 2) // length & capacity
	s2 = append(s2, 1)
	fmt.Println(s2, len(s2), cap(s2)) // [1] 1 2
}
