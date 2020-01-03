package main

import (
    "time"
)

func main() {
    ch := make(chan int)
    go f(ch)
    go p(ch)
    println("Start:")
    time.Sleep(1000 * time.Millisecond)
}

func f(ch chan int) {
    for n := 1; n <= 20; n++ {
        ch <- n
    }
}

func p(ch chan int) {
    for {
        v := <-ch
        println("From channel: ", v)
    }
}

/*
Start:
From channel:  1
From channel:  2
From channel:  3
From channel:  4
From channel:  5
From channel:  6
From channel:  7
From channel:  8
From channel:  9
From channel:  10
From channel:  11
From channel:  12
From channel:  13
From channel:  14
From channel:  15
From channel:  16
From channel:  17
From channel:  18
From channel:  19
From channel:  20
*/
