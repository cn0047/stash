package main

import "fmt"

func main() {
	v := 1
	p := &v
	fmt.Println(p, "|", *p)

	*p = 2
	fmt.Println("New values:", v, *p)
}