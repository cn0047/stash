package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//one()
	two()
}

func two() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	ctx2, _ := context.WithCancel(ctx) // derived context
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx done, error:", ctx.Err())
			return
		case <-ctx2.Done():
			fmt.Println("ctx2 done, error:", ctx2.Err())
			return
		default:
			fmt.Print(".")
		}
	}
}

func one() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx done, error:", ctx.Err())
			return
		default:
			fmt.Print(".")
		}
	}
}
