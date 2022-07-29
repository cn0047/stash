package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := sortStringChars("strings all here") // '  aeeghillnrrsst'

	fmt.Printf("'%s'\n", s)
}

// IMPORTANT: naive algorithm.
func sortStringChars(in string) string {
	s := strings.Split(in, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
