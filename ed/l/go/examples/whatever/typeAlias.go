package main

import (
	"fmt"
)

type V int

type V2 = V

func main() {
	var myvar1 V = 100
	var myvar2 V2 = 200
	fmt.Println(myvar1, myvar2)
}
