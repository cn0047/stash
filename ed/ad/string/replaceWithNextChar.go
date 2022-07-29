package main

import (
	"fmt"
)

func main() {
	// r := Solution("ab?ac?") // abcacd
	r := Solution("re?a?z???") // refabzabc
	// r := Solution([]int{1, 1})
	fmt.Printf("%+v\n", r)
}

func Solution(riddle string) string {
	n := len(riddle)
	s := ""
	p := "a" // previous char

	for i := 0; i < n; i++ {
		c := string(riddle[i])
		if c == "?" {
			tmp := p[0]
			tmp++ // next char (rune value +1)
			if i != n-1 && riddle[i+1] == tmp { // compare with next after current char
				tmp++
			}
			if tmp >= 122 { // z
				tmp = 97 // a
			}
			c = string(tmp)
		}
		s += c
		p = c
	}

	return s
}
