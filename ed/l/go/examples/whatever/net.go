package main

import (
	"fmt"
	"net"
)

func main() {
	f1()
}

func f1() {
	conn, err := net.Dial("tcp", "localhost:11211")
	if err != nil {
		fmt.Printf("🔴 %#v \n", err)
	}
	fmt.Printf("🎾 %#v \n", conn)
}
