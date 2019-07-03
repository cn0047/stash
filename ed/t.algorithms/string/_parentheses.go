package main

import (
	"fmt"
)

func main() {
	fmt.Println(f("(()())(())") == "()()()")
	fmt.Println(f("(()())(())(()(()))") == "()()()()(())")
	fmt.Println(f("(()())(())(()(()))") == "()()()()(())")
	fmt.Println(f("()()") == "")
}

// Remove Outermost Parentheses.
func f(s string) string {
	inStack := 0
	rs := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			inStack++
			if inStack >= 2 {
				rs += string(s[i])
			}
		}
		if s[i] == ')' {
			inStack--
			if inStack >= 1 {
				rs += string(s[i])
			}
		}
	}
	return rs
}
