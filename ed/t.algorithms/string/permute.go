package main

import (
	"fmt"
)

// Print all possible char combinations from given string.
func main() {
	s := "abcd"
	permute(s, 0, len(s)-1)
}

func permute(s string, l int, r int) {
	if l == r {
		fmt.Println(s)
	} else {
		for i := l; i <= r; i++ {
			s = swap(s, l, i)
			permute(s, l+1, r)
			s = swap(s, l, i)
		}
	}
}

func swap(s string, i int, j int) string {
	c1 := s[i]
	c2 := s[j]

	s = s[:i] + string(c2) + s[i+1:]
	s = s[:j] + string(c1) + s[j+1:]

	return s
}
