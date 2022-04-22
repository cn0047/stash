package main

import (
	"fmt"
)

func main() {
	fmt.Println(findWord([]string{"I>N", "A>I", "P>A", "S>P"})) // SPAIN
}

func findWord(arr []string) string {
	// Put into result array first element (characters pair) from input array.
	res := string(arr[0][0]) + string(arr[0][2])
	// Cut off from input array first element.
	arr = arr[1:]

	for len(arr) > 0 {
		s := arr[0] // get next characters pair.
		arr = arr[1:]  // cut off characters pair.
		c1 := s[0]
		c2 := s[2]
		if c2 == res[0] {
			res = string(c1) + res
		} else if c1 == res[len(res)-1] {
			res += string(c2)
		} else {
			arr = append(arr, s) // put back characters pair into input array.
		}
	}

	return res
}
