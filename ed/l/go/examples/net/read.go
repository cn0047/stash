package main

import (
	"io"
	"log"
	"net"
)

// @see:
// in shell #1 go run ed/l/go/examples/net/read.go
// in shell #2 telnet localhost 20080
// in shell #2 ping
func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Failed to bind to port")
	}

	log.Println("Listening on 0.0.0.0:20080")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		log.Println("Received connection")
		go echo(conn)
	}
}

func echo(conn net.Conn) {
	defer conn.Close()

	b := make([]byte, 512)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Printf("Unexpected err: %v \n", err)
			break
		}

		log.Printf("Received %d bytes: %s\n", size, string(b))
		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Failed to write data")
		}
	}
}
