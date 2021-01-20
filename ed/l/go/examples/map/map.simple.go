package main

import (
	"fmt"
)

func main() {
	// f1()
	f2()
}

func f1() {
	m1 := make(map[string]string, 1)
	m1["foo"] = "bar"
	fmt.Println(m1) // map[foo:bar]
}

func f2() {
	lookup := make(map[string]int)
	lookup["goku"] = 9001
	fmt.Printf("Map size: %d \n", len(lookup))
	power, exists := lookup["vegeta"]
	fmt.Println("1:", lookup, power, exists)

	m2 := map[string]string{
		"Bob": "Mr",
		"Amy": "Dr",
	}
	fmt.Println("2:", m2)
	delete(m2, "Amy")
	fmt.Println("3:", m2)
	fmt.Printf("4: %#v\n", m2["J"])

	m5 := map[string]string{}
	m5["a"] = "x"
	m5["b"] = "y"
	fmt.Println("5:", m5)

	m6 := map[string]struct{}{
		"one": {},
		"two": {},
	}
	fmt.Println("6:", m6)
}

/*
Map size: 1
1: map[goku:9001] 0 false
2: map[Bob:Mr Amy:Dr]
3: map[Bob:Mr]
4: ""
5: map[a:x b:y]
6: map[one:{} two:{}]
*/
