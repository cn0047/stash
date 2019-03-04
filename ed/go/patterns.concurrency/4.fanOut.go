package main

import (
  "time"
)

var (
  n = 7
)

func main() {
  chs := []chan int{}

  // init go-routines
  for i := 0; i < n; i++ {
    ch := make(chan int)
    chs = append(chs, ch)
    go gr(i, ch)
  }

  // fan-out to all go-routines
  for _, ch := range chs {
    ch <- 500
  }

  time.Sleep(5 * time.Second)
}

func gr(id int, ch chan int) {
  v := <-ch
  println(id, v)
}
