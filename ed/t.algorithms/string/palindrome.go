package main

import (
	"fmt"
)

func main() {
	fmt.Println(f("this is Palindrome test, looking to: racecar and madam"))
}

func f(str string) []string {
	res := make([]string, 0)

	n := len(str)
	s := ""
	for i := 0; i < n; i++ {
		c := string(str[i])
		if c != " " || i == n-1 {
			s += c
		}
		if c == " " || i == n-1 {
			if IsPalindrome(s) {
				res = append(res, s)
			}
			s = ""
		}
	}

	return res
}

func IsPalindrome(s string) bool {
	n := len(s)

	if n%2 == 0 {
		return false
	}

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true
}
