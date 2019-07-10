package main

import (
	"fmt"
)

func main() {
	m := make(map[int]string)
	m[1] = "a"
	m[2] = "b"
	m[3] = "c"
	fmt.Println(len(m)) // 3
}
