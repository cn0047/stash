package main

import (
	"fmt"
)

var m = map[string]string{"a": "foo", "b": "bar"}

func main() {
	// loopWithKeyAndVal()
	loopWithKeyOnly()
}

func loopWithKeyOnly() {
	for k := range m {
		fmt.Printf("k:%+v; ", k) // k:a; k:b;
	}
}

func loopWithKeyAndVal() {
	for k, v := range m {
		fmt.Printf("k:%+v, v:%+v \n", k, v)
	}
}
