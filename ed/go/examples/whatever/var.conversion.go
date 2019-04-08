package main

import (
	"fmt"
)

type empty interface{}

type example interface {
	notImplemented()
}

func main() {
	one := 1
	var i empty = one
	var e example = i.(example) // panic: interface conversion: int is not main.example: missing method notImplemented
	fmt.Printf("%#v\n", e)
}
