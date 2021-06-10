// @see: https://www.hackerrank.com/challenges/minimum-absolute-difference-in-an-array/problem
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumAbsoluteDifference([]int32{1, -3, 71, 68, 17}))
}

func m(v int32) int32 {
	if v < 0 {
		return v * -1
	}

	return v
}

func minimumAbsoluteDifference(arr []int32) (r int32) {
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	r = m(arr[0] - arr[1])

	if len(arr) == 2 {
		return r
	}

	for i := 1; i < len(arr)-1; i++ {
		j := i + 1
		d := m(arr[i] - arr[j])
		if d < r {
			r = d
		}
	}

	return r
}
