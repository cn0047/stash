package main

import (
  "errors"
  "fmt"
)

var (
  ERR1 = errors.New("ERROR_#1")
  ERR2 = errors.New("ERROR_#2")
)

func main() {
  r := f1()
  fmt.Printf("Main result: %v. \n", r)
}

func f1() int {
  defer Recover([]error{ERR1, ERR2}, func(err interface{}) {
    fmt.Printf("Caught: %v. \n", err.(error))
  })
  f2()
  return 1
}

func f2() {
  panic(ERR2)
}

func Recover(errors []error, cb func(v interface{})) {
  r := recover()

  if r == nil {
    return
  }

  if len(errors) == 0 || inErrors(r, errors) {
    cb(r)
    return
  }

  panic(r)
}

func inErrors(e interface{}, errors []error) bool {
  for _, err := range errors {
    if e == err {
      return true
    }
  }

  return false
}
