package main

import (
	"math/rand"
	"time"
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNum := random(1, 7)
	time.Sleep(time.Second * time.Duration(randomNum))
  slow := ""
  if randomNum > 2 {
    slow = " IT_WAS_SLOW"
  }
  println("Took: ", randomNum, slow)
}
