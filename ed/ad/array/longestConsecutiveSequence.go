// @see: https://leetcode.com/problems/longest-consecutive-sequence
package main

import (
	"fmt"
)

// longestConsecutive returns length for "Longest Consecutive Sequence" for provided nums.
func longestConsecutive(nums []int) (longest int) {
	// Create set of provided numbers.
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}

	for _, v := range nums {
		prevVal := v - 1
		if !set[prevVal] { // it means start of new sequence if there is no prevVal for current value.
			i := 0
			ok := true
			for ok {
				nextVal := v + i
				_, inSet := set[nextVal]
				if inSet {
					i++
				} else {
					ok = false // end of sequence.
				}
			}

			// Determine longest.
			if i > longest {
				longest = i
			}
		}
	}

	return longest
}

func main() {
	r := 0
	r = longestConsecutive([]int{100, 4, 200, 1, 3, 2}) // 4
	// r = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}) // 9
	fmt.Printf("\n===\n%v\n", r)
}
