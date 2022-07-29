package main

import (
	"fmt"
	"sort"
)

func main() {
	var r int
	// r = solution([]int{1, 3, -3})
	// r = solution([]int{4, 3, 2, 5, 1, 1})
	r = solution([]int{-4, -3, -7, -1})
	fmt.Printf("%v \n", r)
}

func modulo(n int) int {
	if n < 0 {
		return n * -1
	}

	return n
}

// solution finds maxDifferenceInArrays.
// IMPORTANT: naive algorithm, O(n log n)!
func solution(arr []int) int {
	sort.Ints(arr)
	n := len(arr)-1
	return modulo(arr[n]-arr[0])
}
