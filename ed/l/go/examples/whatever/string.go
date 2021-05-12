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
	// strBuf()
	subStr()
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

func subStr() {
	s := "abcdefg"
	l := len(s)
	fmt.Printf("subStr: '%s' \n", s[0:])    // subStr: 'abcdefg'
	fmt.Printf("subStr: '%s' \n", s[0:0])   // subStr: ''
	fmt.Printf("subStr: '%s' \n", s[2:2])   // subStr: ''
	fmt.Printf("subStr: '%s' \n", s[l-2:l]) // subStr: 'fg'
	// fmt.Printf("subStr: '%s' \n", s[l-2:l+2]) // panic: runtime error: slice bounds out of range [:9] with length 7
	// fmt.Printf("subStr: '%s' \n", s[6:3])     // error: invalid slice index: 6 > 3
}
