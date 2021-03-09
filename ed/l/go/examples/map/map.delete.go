package main

import (
	"fmt"
)

func main() {
	one()
}

func one() {
	m := map[string]string{
		"foo": "1",
	}
	delete(m, "foo") // ok
	delete(m, "foo") // ok

	fmt.Printf("%v", m)
}
