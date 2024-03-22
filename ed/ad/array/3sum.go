// @see: https://leetcode.com/problems/3sum
package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int
	// nums = []int{0, 0, 0} // [[0,0,0]]
	// nums = []int{0, 1, 1} // [[0,0,0]]
	// nums = []int{-1, 0, 1, 2, -1, -4} // [[-1,-1,2],[-1,0,1]]
	nums = []int{-2, -2, 0, 0, 2, 2} // [[-2 0 2]]
	// nums = []int{-1, 0, 2, 1, 2, -1, -4, 0} // probably: [[-4 2 2],[-1 -1 2],[-1 0 1]]
	// nums = []int{-3, 3, 4, 2, 0, -3, 1, 2, -4} //
	r := threeSum(nums)
	fmt.Printf("Result:\n%v \n", r)
}

// threeSum returns array of elements sum of which equals to 0.
// plan: sort; point to 1st element + 2sum;
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)

	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			}
		}

		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}
