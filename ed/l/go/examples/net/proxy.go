package main

import (
	"io"
	"log"
	"net"
)

// @see: curl -i -X GET localhost:80
func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("failed to bind to port")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("failed to accept connection")
		}

		go handle(conn)
	}
}

func handle(src net.Conn) {
	dst, err := net.Dial("tcp", "x.website:80")
	if err != nil {
		log.Fatalln("failed to connect to unreachable host")
	}
	defer dst.Close()

	go func() {
		// Transmit request.
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	// Transmit response.
	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}
