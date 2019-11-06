package main

import (
	"fmt"
)

type G struct {
	x string
}

type P struct {
	k string
	v int
	G
}

func main() {
	p := P{k: "1", v: 2, G: G{x: "y"}}
	fmt.Printf("%v", p)
}
