package main

import (
	"fmt"
)

var (
	n = 0
)

func main() {
	for i := 0; i < 15; i++ {
		go func(i int) {
			println(i)
			n++ // here
		}(i)
	}
	for {
		if n == 5 {
			break
		}
	}
	fmt.Scanln()
}

/*
1
3
4
0
2
*/
