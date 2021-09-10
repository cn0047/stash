package main

import (
	"fmt"
	"net"
	"os"
)

// @see:
// go run ed/l/go/examples/net/dial.go dial "scanme.nmap.org:80"
// go run ed/l/go/examples/net/dial.go scan "scanme.nmap.org"
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Provide command.\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "dial":
		dial()
	case "scan":
		scan()
	}
}

func scan() {
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			addr := fmt.Sprintf("%s:%d", os.Args[2], j)
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				fmt.Printf("failed to dial: %s, err: %v \n", addr, err)
				return
			}
			defer conn.Close()
			fmt.Printf("Connection to: %s successful.\n", addr)
		}(i)
	}
	fmt.Scanln()
}

func dial() {
	if len(os.Args) < 3 {
		fmt.Printf("Provide address.\n")
		os.Exit(1)
	}

	addr := os.Args[2]
	conn, err := net.Dial("tcp", addr)
	if err == nil {
		fmt.Printf("Connection to: %s successful.\n", addr)
	}
	defer conn.Close()
}
