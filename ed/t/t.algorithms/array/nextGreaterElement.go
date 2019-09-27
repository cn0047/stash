package main

import (
	"fmt"
)

func main() {
	fmt.Println(f([]int{1, 4, 5, 9, 7}))
	fmt.Println(f([]int{-1, 4, 15, 9, 7}))
	fmt.Println(f([]int{-1, 0}))
	fmt.Println(f([]int{9, 9}))
	fmt.Println(f([]int{3, 7, 9}))
}

// f returns next greater element from given array.
// This solution based on stack usage.
func f(a []int) int {
	var stack []int
	if a[0] > a[1] {
		stack = []int{a[0], a[1]}
	} else {
		stack = []int{a[1], a[0]}
	}

	for i := 2; i < len(a); i++ {
		if a[i] > stack[0] {
			stack[1] = stack[0]
			stack[0] = a[i]
		} else if a[i] > stack[1] {
			stack[1] = a[i]
		}
	}

	return stack[1]
}
