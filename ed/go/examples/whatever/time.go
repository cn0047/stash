package main

import (
	"fmt"
	"time"
)

func main() {
  f1()
	f2()
}

func f1() {
  fmt.Printf("%+v ↔️ %+v \n", time.Now(), time.Now().Add(-24*time.Hour))
}

func f2() {
  dateFormat := "2006-01-02"
  d1, _ := time.Parse(dateFormat, "2018-08-01")
  d2, _ := time.Parse(dateFormat, "2028-08-01")
  fmt.Printf(
    "2️⃣ d1: %+v | d2: %+v | %+v, %+v \n",
    d1,
    d2,
    d1.After(time.Now()),
    d2.After(time.Now()),
  )
}
