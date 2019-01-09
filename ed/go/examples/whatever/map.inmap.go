package main

import "fmt"

func main() {
	m := make(map[string]map[string]string)
	m["a"] = make(map[string]string)
	m["a"]["x"] = "1"
	m["a"]["y"] = "2"
	m["a"]["z"] = "2"
	fmt.Printf("%+v\n", m)
	fmt.Printf("%+v\n", m["a"]["z"])
}
