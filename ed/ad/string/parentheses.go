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
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			inStack++
			if inStack >= 2 {
				res += string(s[i])
			}
		}
		if s[i] == ')' {
			inStack--
			if inStack >= 1 {
				res += string(s[i])
			}
		}
	}

	return res
}
