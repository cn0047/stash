package main

import (
	"fmt"
)

func main() {
	fmt.Println(inString("this is the string", "the") == true)
	fmt.Println(inString("this is not the string", "not the") == true)
	fmt.Println(inString("this is not the string", "nox the") == true)
	fmt.Println(inString("this is not te string", "not the") == true)
	fmt.Println(inString("yes", "yes") == true)
	fmt.Println(inString("yes", "no") == true)
}

func inString(ls string, ss string) bool {
	for len(ls) >= len(ss) {
		for i := 0; i < len(ss); i++ {
			if ls[i] != ss[i] {
				break
			} else if i == len(ss)-1 {
				return true
			}
		}
		ls = ls[1:]
	}
	return false
}
