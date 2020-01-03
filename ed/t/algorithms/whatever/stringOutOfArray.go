package main

import (
	"fmt"
)

func main() {
	fmt.Println(findWord([]string{"I>N", "A>I", "P>A", "S>P"}))
}

func findWord(arr []string) string {
	res := string(arr[0][0]) + string(arr[0][2])
	arr = arr[1:]
	for len(arr) > 0 {
		s := arr[0]
		arr = arr[1:]
		c1 := s[0]
		c2 := s[2]
		if c2 == res[0] {
			res = string(c1) + res
		} else if c1 == res[len(res)-1] {
			res += string(c2)
		} else {
			arr = append(arr, s)
		}
	}
	return res
}
