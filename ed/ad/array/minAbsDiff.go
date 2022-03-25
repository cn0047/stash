// @see: https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsoluteDifference([]int32{1, -3, 71, 68, 17}))
}

// minimumAbsoluteDifference returns minimum absolute difference in an array.
// The absolute difference is the positive difference between two values a and b.
func minimumAbsoluteDifference(arr []int32) (diff int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	diff = modulo(arr[0] - arr[1])

	if len(arr) == 2 {
		return diff
	}

	for i := 1; i < len(arr)-1; i++ {
		j := i + 1
		currentDiff := modulo(arr[i] - arr[j])
		if currentDiff < diff {
			diff = currentDiff
		}
	}

	return diff
}

func modulo(v int32) int32 {
	if v < 0 {
		return v * -1
	}

	return v
}
