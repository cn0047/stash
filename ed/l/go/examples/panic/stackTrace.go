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

/*
[RECOVER] panic | goroutine 1 [running]:
runtime/debug.Stack(0xc00007ae48, 0x10ace20, 0x10eb820)
*/
