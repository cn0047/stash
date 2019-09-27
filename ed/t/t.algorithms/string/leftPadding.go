package main

import (
	"fmt"
)

func main() {
	fmt.Println(LeftPadding("yes", "!", 4) == "!yes")
	fmt.Println(LeftPadding("yes", "!", 3) == "yes")
	fmt.Println(LeftPadding("yes", "!", 7) == "!!!!yes")
}

func LeftPadding(s string, x string, n int) string {
	if len(s) >= n {
		return s
	}

	m := n - len(s)
	res := ""
	for i := 0; i < m; i++ {
		res += x
	}
	res += s

	return res
}
