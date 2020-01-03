package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	default:
		fmt.Println("No message received.")
	}

	select {
	case messages <- "hey":
		fmt.Println("Sent message")
	default:
		fmt.Println("No message sent.")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message: ", msg)
	default:
		fmt.Println("No message received.")
	}
}

/*
No message received.
No message sent.
No message received.
*/
