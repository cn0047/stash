package main

import (
  "fmt"
)

var (
  LogPanic = false
)

func main() {
  r := f1()
  println(r)
  fmt.Scanln()
}

func f1() (r int) {
  defer func() {
    p := recover()
    if p != nil {
      if LogPanic { // current panic is not LogPanic
        println("[panic] ", p)
      }
      r = -1
    }
  }()

  res := panicFunc()
  r = res

  return r
}

func panicFunc() (r int) {
  panic("panic")
}

/*
Result:
-1
*/
