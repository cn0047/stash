// kmp - Knuth-Morris-Pratt string matching algorithm.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(kmp("aaaabaaaaabbbaaaab", "aaab") == 1)
	fmt.Println(kmp("thisx isy strisg", "isg") == 13)
	fmt.Println(kmp("abcabd abcabx", "abcabx") == 7)
	fmt.Println(kmp("well, it works", "ork") == 10)
	fmt.Println(kmp("well, it works", "orx") == -1)
	fmt.Println(kmp("knp", "knp") == 0)

}

func kmp(s string, p string) int {
	i := 0
	n := len(s)
	j := 0
	m := len(p)

	t := getPatternTable(p)

	for i < n {
		if s[i] == p[j] {
			i++
			j++
			if j == m {
				return i - j
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = t[j-1]
			}
		}
	}

	return -1
}

func kmp0(s string, p string) int {
	i := 0
	n := len(s)
	j := 0
	m := len(p)

	t := getPatternTable(p)

	for i < n {
		if s[i] == p[j] {
			i++
			j++
		}
		if j == m {
			return i - j
		} else if i < n && s[i] != p[j] {
			if j == 0 {
				i++
			} else {
				j = t[j-1]
			}
		}
	}

	return -1
}

// getPatternTable generates table with values like this:
// input string: a b c d a b e a b f
// array values: 0 0 0 0 1 2 0 1 2 0
func getPatternTable(p string) []int {
	n := len(p)
	t := make([]int, n)

	j := 0
	for i := 1; i < n; i++ {
		if p[i] == p[j] {
			j++
			t[i] = j
		} else {
			if j == 0 {
				t[i] = 0
			} else {
				j = t[j-1]
				i--
			}
		}
	}

	return t
}
