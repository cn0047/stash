package main

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed desc.txt
	desc string
)

func main() {
	fmt.Printf("desc: %s", desc)
}
