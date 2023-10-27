package main

import (
	"fmt"
)

func main() {
	// f1()
	f2()
	// swap()
}

func swap() {
	s := []string{"a", "b", "c"}
	s[1], s[2] = s[2], s[1]
	fmt.Println("[swap]", s) // [swap] [a c b]
}

func f2() {
	s := []string{"a", "b", "c", "d"}
	fmt.Println("[f2]", s[:1])    // [f2] [a]
	fmt.Println("[f2]", s[1:])    // [f2] [b c d]
	fmt.Println("[f2]", s[0:2:2]) // [f2] [a b]
	fmt.Println("[f2]", s[0:3:4]) // [f2] [a b c]
}

// f1 describes slice length & capacity.
func f1() {
	s1 := make([]int, 2) // length
	s1 = append(s1, 1)
	fmt.Println("[f1]", s1, len(s1), cap(s1)) // [f1] [0 0 1] 3 4

	s2 := make([]int, 0, 2) // length & capacity
	s2 = append(s2, 1)
	fmt.Println("[f1]", s2, len(s2), cap(s2)) // [f1] [1] 1 2
}
