// Longest String Chain.
// @see: https://leetcode.com/problems/longest-string-chain
package main

import (
	"fmt"
	"sort"
)

func main() {
	var r int
	r = f([]string{"a", "b", "ba", "bca", "bda", "bdca"}) // must be 4
	// r = f([]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"})
	// r = f([]string{"abcd","dbqca"})
	// r = f([]string{"a","aa","aab","aabb","bbaac"})
	// r = f([]string{"a","ab","ac","bd","abc","abd","abdd"}) // must be 4
	// r = f([]string{"i", "in", "sin", "sing", "sign", "sting", "signs", "string", "strings"})
	fmt.Printf("%v \n", r)
}

func f(input []string) int {
	sort.Slice(input, func(i, j int) bool { return len(input[i]) < len(input[j]) })

	res := 0
	cache := map[string]int{}

	for i := 0; i < len(input); i++ {
		word := input[i]

		longestPredecessor := 0
		// Make subword by deleting 1 char from word,
		// find in cache longest predecessor for subword.
		for j := 0; j < len(word); j++ {
			subword := word[0:j] + word[j+1:]
			if v, ok := cache[subword]; ok {
				longestPredecessor = max(longestPredecessor, v)
			}
		}
		longestPredecessor++
		cache[word] = longestPredecessor
		res = max(longestPredecessor, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
