// @see: https://www.hackerrank.com/challenges/ctci-ice-cream-parlor/problem
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findSumInArrayNoSort([]int32{1, 4, 5, 3, 2}, 4))
	fmt.Println(findSumInArrayNoSort([]int32{-2, 1, 2, 4, 5, 2, 3}, 4))
	fmt.Println(findSumInArrayNoSort([]int32{1, 4, 5, -3, 2, 13, 0, 4, 0, 3, -3}, 4))
}

// findSumInArray returns 2 values from array arr
// sum of which will be equal to seaking sum value.
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

// findSumInArray returns 2 indexes to elements in array arr
// sum of which will be equal to seaking sum value.
func findSumInArrayNoSort(arr []int32, sum int32) (int32, int32) {
	value := arr[0]
	index := int32(0)
	cache := map[int32]int32{value: index}

	for i := 1; i < len(arr); i++ {
		value := arr[i]
		diff := sum - value
		index, ok := cache[diff]
		if ok {
			return index, int32(i)
		}
		cache[value] = int32(i)
	}

	return -1, -1
}
