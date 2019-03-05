package main

import (
  "sync"
  "time"
)

func main()  {
  wg := sync.WaitGroup{}
  wg.Add(2)
  go func() {
    for {
      println(1)
      time.Sleep(1*time.Second)
    }
  }()
  go func() {
    for {
      println(2)
      time.Sleep(1*time.Second)
    }
  }()
  go func() {
    for {
      println(3)
      time.Sleep(1*time.Second)
    }
  }()
  wg.Wait()
}

// in `ps aux` you'll see only 1 record
