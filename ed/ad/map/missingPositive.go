// @see *: https://leetcode.com/problems/first-missing-positive
package main

import (
	"fmt"
)

func main() {
	r := 0
	r = firstMissingPositive([]int{1, 2, 0})           // 3
	r = firstMissingPositive([]int{3, 4, -1, 1})       // 2
	r = firstMissingPositive([]int{7, 8, 9, 11, 12})   // 1
	r = firstMissingPositive([]int{-7, 8, 9, -11, 12}) // 1
	r = firstMissingPositive([]int{1, 9, 100, 3, 57})  // 2
	r = firstMissingPositive([]int{1, 30, 2, 15, 4})   // 3
	fmt.Printf("===\n%v\n", r)
}

func firstMissingPositive(nums []int) int {
	n := len(nums)

	maxPossible := n + 1 // Max possible value in array with given length.
	// Replace negative values and values with way-way bigger possible value to maxPossible.
	for i := 0; i < n; i++ {
		if nums[i] <= 0 || nums[i] > n {
			nums[i] = maxPossible
		}
	}

	// Use "-" sign as flag for position in array which means value exists somewhere in entire array.
	for i := 0; i < n; i++ {
		val := abs(nums[i])
		if val <= n {
			nums[val-1] = -abs(nums[val-1])
		}
	}

	// Only "-" sign matters here.
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
