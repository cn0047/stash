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
