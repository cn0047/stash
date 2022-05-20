package main

import (
	"fmt"
)

func main() {
	r := rotateArrayToTheRight([]int{1,2,3}, 5)
	fmt.Printf("[result] %v \n", r) // [2 3 1]
}

// rotateArrayToTheRight rotates array to the right rotationCount times.
func rotateArrayToTheRight(a []int, rotationCount int) []int {
	l := len(a)
	arr := make([]int, l, l)

	// No need to rotate array, just put element into right place.
	for i:=0; i<l; i++ {
		idx := (i+rotationCount)%l
		arr[idx] = a[i]
	}

	return arr
}
