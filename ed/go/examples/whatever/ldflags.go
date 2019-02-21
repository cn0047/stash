package main

import (
  "fmt"
)

var (
  CommitHash string
)

func main() {
  fmt.Printf("CommitHash = %v \n", CommitHash)
}

// go build -ldflags "-X main.CommitHash="`git rev-parse HEAD` -o ldflags ed/go/examples/whatever/ldflags.go
