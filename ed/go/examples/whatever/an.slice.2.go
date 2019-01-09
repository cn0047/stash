package main

import (
	"fmt"
)

func main() {
	fs1()
	slice3()
	three()
	three2()
}

func fs1() {
	a := make([]int, 0)
	a = append(a, 1, 2, 3)  // few elements
	fmt.Printf("%+v \n", a) // [1 2 3]

	b := []int{4, 5}
	a = append(a, b...)     // concatenation
	fmt.Printf("%+v \n", a) // [1 2 3 4 5]
}

func slice3() {
	a := []int{0, 1, 2, 3, 4, 5}
	s := a[2:3:4]           // from:to:capacity
	fmt.Printf("%+v \n", s) // [2]
}

func three() {
	fmt.Println("\n three: -------------------")
	a := [...]string{"r", "o", "a", "d"}
	a2 := a[:]
	a[3] = "x"
	fmt.Printf("a = %+v, a2 = %+v \n", a, a2) // a = [r o a x], a2 = [r o a x]

	s := []string{"r", "o", "a", "d"}
	s2 := s[:]
	s2[3] = "x"
	fmt.Printf("s = %+v, s2 = %+v \n", s, s2) // s = [r o a x], s2 = [r o a x]

	c := []string{"a"}
	c2 := make([]string, 1)
	copy(c2, c)
	c2[0] = "b"
	fmt.Printf("c = %+v, c2 = %+v \n", c, c2) // c = [a], c2 = [b]

}

func three2() {
	fmt.Println("\n three2: -------------------")

	s := make([]int, 3, 5)
  // s[5] = 1 // panic: runtime error: index out of range
  _ = s

	// s2 := make([]int, 5, 3) // len larger than cap in make([]int)
}
