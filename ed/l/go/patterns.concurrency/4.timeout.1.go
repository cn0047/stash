package main

import (
  "fmt"
  "time"
)

func main() {
  timeout := make(chan bool, 1)
  ch := make(chan int, 1)

  go func() {
      time.Sleep(1 * time.Second)
      timeout <- true
  }()

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
