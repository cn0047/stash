package main

import (
	"fmt"
)

func main() {
	s := make([]int, 1)
	f(s)
	fmt.Println("end:", s) // end: [0]
}

func f(s []int) {
	s = append(s, 1)
	fmt.Println("in f:", s) // in f: [0 1]
}
