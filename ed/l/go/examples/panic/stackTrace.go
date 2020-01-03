package main

import (
  "fmt"
  "runtime/debug"
)

func main() {
  defer func() {
    if r := recover(); r != nil {
      stack := debug.Stack()
      str := string(stack[:])
      fmt.Printf("[RECOVER] %+v | %s \n", r, str)
    }
  }()
  f1()
}

func f1() {
  f2()
}

func f2() {
  panic("panic")
}
