// @see: https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findSumInArray([]int32{1, 4, 5, 3, 2}, 4))
}

func findSumInArray(arr []int32, sum int32) (int32, int32) {
	n := len(arr)

	sarr := make([]int32, n)
	copy(sarr, arr)
	sort.Slice(sarr, func(i, j int) bool { return sarr[i] < sarr[j] })

	l := 0
	r := n - 1
	for l < r {
		if sarr[l]+sarr[r] > sum {
			r--
		} else if sarr[l]+sarr[r] < sum {
			l++
		} else {
			return sarr[l], sarr[r]
		}
	}

	return -1, -1
}
