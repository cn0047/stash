// @see: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string
// @see: https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii

package main

import (
	"fmt"
	"strings"
)

func f(s string, k int) string {
	// stack holds slices where each slice is:
	// pair of int representation for char and count for this char.
	stack := [][]int{}

	for i := 0; i < len(s); i++ {
		char := int(s[i])
		if len(stack) > 0 && stack[len(stack)-1][0] == char {
			stack[len(stack)-1][1]++
		} else {
			stack = append(stack, []int{char, 1})
		}

		if stack[len(stack)-1][1] == k {
			stack = stack[:len(stack)-1]
		}
	}

	var res strings.Builder
	for i := 0; i < len(stack); i++ {
		char := rune(stack[i][0])
		count := stack[i][1]
		res.WriteString(strings.Repeat(string(char), count))
	}

	return res.String()
}

func main() {
	s, k := "", 0
	s, k = "abcd", 2 // abcd
	s, k = "deeedbbcccbdaa", 3 // aa
	s, k = "pbbcggttciiippooaais", 2 // ps
	s, k = "abbaca", 2               // ca
	s, k = "azxxzy", 2               // ay
	r := f(s, k)
	fmt.Printf("===\n%v\n", r)
}
