package main

import (
	"fmt"
)

func main() {
	f1()
	f2()
}

func f1() {
	var a = []string{"1"}
	uparr(a)
	fmt.Printf("%+v\n", a) // [2]
}

func uparr(a []string) {
	a[0] = "2"
}

func f2() {
	m := make(map[string]string)
	upmap(m)
	fmt.Printf("%+v\n", m) // map[a:1]
}

func upmap(m map[string]string) {
	m["a"] = "1"
}
