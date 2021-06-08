package main

import (
	"fmt"
)

func main() {
	deleteInLoop()
	// deleteTwice()
}

func deleteInLoop() {
	m := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
	}
	for k, v := range m {
		if v%2 == 0 {
			delete(m, k)
		}
	}

	fmt.Printf("[deleteInLoop] %v \n", m) //  map[a:1 c:3 e:5]
}

func deleteTwice() {
	m := map[string]string{
		"foo": "1",
		"bar": "2",
	}
	delete(m, "foo") // ok
	delete(m, "foo") // ok

	fmt.Printf("[deleteTwice] %v \n", m) // map[bar:2]
}
