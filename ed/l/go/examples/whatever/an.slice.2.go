package main

import (
	"fmt"
)

func main() {
	//concatenation()
	//slice3()
	//copyByLink()
	copySimple()
	//four()
}

func concatenation() {
	a := make([]int, 0)
	a = append(a, 1, 2, 3)
	fmt.Printf("%+v \n", a) // [1 2 3]

	b := []int{4, 5}
	a = append(a, b...)     // concatenation
	fmt.Printf("%+v \n", a) // [1 2 3 4 5]
}

func slice3() {
	a := []int{0, 1, 2, 3, 4, 5}
	s := a[2:3:4]                    // from:to:capacity
	fmt.Printf("[slice3] %+v \n", s) // [2]
}

// copyByLink describes copy by link cases.
func copyByLink() {
	a := [...]string{"r", "o", "a", "d"} // array
	a2 := a[:]
	a[3] = "x"
	fmt.Printf("[copyAllegedly] a = %+v, a2 = %+v \n", a, a2) // a = [r o a x], a2 = [r o a x]

	s := []string{"r", "o", "a", "d"} // slice
	s2 := s[:]
	s2[3] = "x"
	fmt.Printf("[copyAllegedly] s = %+v, s2 = %+v \n", s, s2) // s = [r o a x], s2 = [r o a x]
}

func copySimple() {
	c1 := []string{"a"}
	c2 := make([]string, 1)
	r1 := copy(c2, c1)
	c2[0] = "b"
	fmt.Printf("[copySimple] c1: %+v, c2: %+v, r: %+v \n", c1, c2, r1) // c1: [a], c2: [b], r: 1

	s1 := []string{"1", "2", "3"}
	s2 := []string{"4", "5"}
	r2 := copy(s1, s2)
	fmt.Printf("[copySimple] s1: %+v, r: %+v \n", s1, r2) // s1: [4 5 3], r: 2

	a1 := []string{"1", "2"}
	a2 := []string{"3", "4", "5"}
	r3 := copy(a1, a2)
	fmt.Printf("[copySimple] a1: %+v, r: %+v \n", a1, r3) // a1: [3 4], r: 2
}

func four() {
	s := make([]int, 3, 5)
	// s[5] = 1 // panic: runtime error: index out of range
	_ = s

	// s2 := make([]int, 5, 3) // len larger than cap in make([]int)
}
