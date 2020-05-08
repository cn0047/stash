package main

import (
	"fmt"
)

func main() {
	// f1()
	f2()
}

func f2() {
	v := 1
	changeValue(&v)
	fmt.Println("[2]", v) // [2] 2
}

func changeValue(p *int) {
	*p = 2
}

func f1() {
	v := 1
	p := &v
	fmt.Println("[1]", p, "|", *p) // [1] 0xc000098008 | 1

	*p = 2
	fmt.Println("[1] New values:", v, *p) // [1] New values: 2 2
}
