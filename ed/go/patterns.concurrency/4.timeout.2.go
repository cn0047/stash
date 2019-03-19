package main

import (
  "fmt"
  "time"
)

func main() {
  timeout := time.After(1 * time.Second)
  ch := make(chan int, 1)

  go func() {
    time.Sleep(2 * time.Second)
    ch <- 204
  }()

  select {
  case v := <-ch:
    fmt.Println("received result:", v)
  case <-timeout:
    fmt.Println("timeout")
  }
}
