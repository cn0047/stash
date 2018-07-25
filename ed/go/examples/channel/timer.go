package main

import "time"

func main() {
	done := time.After(500 * time.Millisecond)
	for {
		select {
		case <-done:
			print("done")
			return
		default:
			print(".")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
