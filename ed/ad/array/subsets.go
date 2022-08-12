package main

import (
	"fmt"
)

func main() {
	r := subsets([]int{1, 2, 3})
	fmt.Printf("%v \n", r)
}

// subsets generates all possible subsets for given slice nums.
// Example: iput: [1,2,3], output: [[] [1] [2] [2 1] [3] [3 1] [3 2] [3 2 1]]
// @see: https://leetcode.com/problems/subsets/
func subsets(nums []int) [][]int {
	res := [][]int{{}}

	for i := 0; i < len(nums); i++ {
		el1 := nums[i]
		tmp := [][]int{{el1}}
		for j := 1; j < len(res); j++ {
			el2 := res[j]
			x := append([]int{el1}, el2...)
			tmp = append(tmp, x)
		}
		res = append(res, tmp...)
	}

	return res
}
