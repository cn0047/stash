package main

import (
  "fmt"
  "time"
)

const (
  timeoutThresholdInSeconds = 1
  resultThresholdInSeconds = 2
)

func main() {
  timeout := make(chan bool, 1)
  ch := make(chan int, 1)

  go func() {
      time.Sleep(time.Second * timeoutThresholdInSeconds)
      timeout <- true
  }()

  go func() {
    time.Sleep(time.Second * resultThresholdInSeconds)
    ch <- 204
  }()

  select {
  case v := <-ch:
      fmt.Println("received result:", v)
  case <-timeout:
      fmt.Println("timeout")
  }
}
