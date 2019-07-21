package main

import (
	"strings"
)

func main() {
	println(uCaseFirst("test"))
	println(uCaseFirst("x test"))
}

func uCaseFirst(s string) string {
	return strings.Title(s[0:1]) + s[1:]
}
