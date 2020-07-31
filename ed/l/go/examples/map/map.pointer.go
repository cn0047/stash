package main

import (
	"fmt"
)

func main() {
	f1()
}

func f1() {
	m1 := make(map[string]string, 1)
	m1["foo"] = "bar"
	fmt.Println(m1)
}
