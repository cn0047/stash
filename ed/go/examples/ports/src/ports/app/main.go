package main

import (
	"fmt"
	"os"

	"ports/app/rr"
)

// todo: add os signals handling.
func main() {
	if len(os.Args) < 1 {
		fmt.Println("you have to provide additional param: client or server")
		return
	}

	param := os.Args[1]
	switch param {
	case "client":
		rr.StartClient()
	case "server":
		rr.StartServer()
	default:
		fmt.Printf("got unsupported param: %s, but wants: client or server\n", param)
	}
}
