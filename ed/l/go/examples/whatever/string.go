package main

import (
	"bytes"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // (magic)
	return fmt.Sprintf("\n%v (%v years)", p.Name, p.Age)
}

func main() {
	// printStruct()
	strBuf()
}

func printStruct() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

func strBuf() {
	buf := bytes.NewBufferString("my string")
	fmt.Printf("buf: %s \n", buf)
}
