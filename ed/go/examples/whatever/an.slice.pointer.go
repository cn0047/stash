package main

import (
	"fmt"
)

func main() {
	s := []int{1}
	fmt.Println(s) // [1]
	f(&s)
	fmt.Println(s) // [1 2]
}

func f(a *[]int) {
	*a = append(*a, 2)
}
