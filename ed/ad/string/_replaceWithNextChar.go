package main

import (
	"fmt"
)

func main() {
	r := Solution("ab?ac?")
	// r := Solution("re?a?z???")
	r := Solution([]int{1, 1})
	fmt.Println("---")
	fmt.Printf("%+v\n", r)
}

func Solution(riddle string) string {
	n := len(riddle)
	s := ""
	p := "a"
	for i := 0; i < n; i++ {
		c := string(riddle[i])
		if c == "?" {
			tmp := p[0]
			tmp++
			if i != n-1 && riddle[i+1] == tmp {
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
