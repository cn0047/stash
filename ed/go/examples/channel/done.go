package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go worker(done)
	// Important to obtain value from channel (to stop script).
	v := <-done
	fmt.Printf("%+v\n", v)
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Value passed into channel not important here.
	done <- false
}

/*
working...done
false
*/
