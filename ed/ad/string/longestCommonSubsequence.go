// @see: https://leetcode.com/problems/longest-common-subsequence
package main

import (
	"fmt"
)

// lcs represents Longest Common Subsequence.
func lcs(str1, str2 string) (string, int) {
	len1, len2 := len(str1), len(str2)

	// Init cache table.
	cache := make([][]int, len1+1)
	for i := range cache {
		cache[i] = make([]int, len2+1)
	}
	// Fill cache table.
	for i := 0; i <= len1; i++ {
		for j := 0; j <= len2; j++ {
			if i == 0 || j == 0 {
				cache[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				cache[i][j] = cache[i-1][j-1] + 1
			} else {
				if cache[i-1][j] > cache[i][j-1] {
					cache[i][j] = cache[i-1][j]
				} else {
					cache[i][j] = cache[i][j-1]
				}
			}
		}
	}

	// Get LCS from cache.
	l := cache[len1][len2]
	res := make([]byte, l)
	a := len1
	b := len2
	for a > 0 && b > 0 {
		if str1[a-1] == str2[b-1] {
			res[l-1] = str1[a-1]
			a--
			b--
			l--
		} else if cache[a-1][b] > cache[a][b-1] {
			a--
		} else {
			b--
		}
	}

	return string(res), cache[len1][len2]
}

func main() {
	a, b := "", ""
	a, b = "abc", "xyz"       // ""
	a, b = "absdhs", "abdhsp" // "abdhs"
	a, b = "abcde", "ace"     // "ace"
	a, b = "abc", "abc"       // "abc"
	a, b = "stone", "longest" // "one"

	res, l := lcs(a, b)
	fmt.Printf("res: %v %s \n", l, res)
}
