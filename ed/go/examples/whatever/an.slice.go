package main

import "fmt"

func main() {
	zero()
	two()
	one()
}

func zero() {
	fmt.Println("\n zero: -------------------")
	letters := []string{"a", "b", "c", "d"}
	fmt.Printf("letters: %+v \n", letters) // letters: [a b c d]

	scores := make([]int, 10)     // slice
	fmt.Printf("scores: %+v \n", scores) // scores: [0 0 0 0 0 0 0 0 0 0]

	scores2 := make([]int, 0, 10) // slice of length 0 and capacity 10.
	fmt.Printf("scores2: %+v \n", scores2) // scores2: []

	scores3 := make([]int, 0, 10)
	//scores3[7] = 903 // panic: runtime error: index out of range
	fmt.Printf("scores3: %+v \n", scores3) // scores3: []

	a1 := []string{"a", "b"}
	a2 := []string{"c", "d", "e"}
	a3 := append(a1, a2...)
	fmt.Printf("as3: %+v \n", a3) // as3: [a b c d e]
}

func two() {
	fmt.Println("\n two: -------------------")
	// Unlike an array type, a slice type has no specified length.
	// A slice literal is declared just like an array literal, except you leave out the element count.
	var s1 []byte // it's slice because length not specified
	fmt.Printf("s1 ❶: %+v, len = %d, cap = %d \n", s1, len(s1), cap(s1)) // s1 ❶: [], len = 0, cap = 0

	s1 = append(s1, 3)
	fmt.Printf("s1: %+v \n", s1) // s1: [3]

	s1Add := s1[:]
	fmt.Printf("s1Add: %+v \n", s1Add) // s1Add: [3]

	s2 := []byte{0, 0, 0, 0, 0}
	fmt.Printf("s2: %+v, len = %d, cap = %d \n", s2, len(s2), cap(s2)) // s2: [0 0 0 0 0], len = 5, cap = 5

}

func one() {
	fmt.Println("\n one: -------------------")
	s := make([]int, 0)
	s = append(s, 27)
	fmt.Printf("\n %#v", s) // []int{27}

	s2 := mutateSlice(s)
	fmt.Printf("\n s  = %#v", s) // s  = []int{27}
	fmt.Printf("\n s2 = %#v", s2) // s2 = []int{27, 27}
}

func mutateSlice(s []int) []int {
	s = append(s, 27)
	return s
}
