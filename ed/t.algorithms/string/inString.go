package main

import (
	"fmt"
)

func main() {
	fmt.Println(inString("this is the string", "the") == true)
	fmt.Println(inString("this is not the string", "not the") == true)
	fmt.Println(inString("this is not the string", "nox the") == false)
	fmt.Println(inString("this is not te string", "not the") == false)
	fmt.Println(inString("yes", "yes") == true)
	fmt.Println(inString("yes", "no") == false)
	fmt.Println(inString("this is not barfoo", "foo") == true)
}

func inString0(ls string, ss string) bool {
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

func inString(inStr string, subStr string) bool {
	for j := 0; j <= len(inStr)-len(subStr); j++ {
		for i := 0; i < len(subStr); i++ {
			if inStr[j+i] == subStr[i] {
				if i == len(subStr)-1 {
					return true
				}
			} else {
				break
			}
		}
	}
	return false
}
