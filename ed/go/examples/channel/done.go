package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    // Value passed into channel not important here.
    done <- false
}

func main() {
    done := make(chan bool, 1)
    go worker(done)
    // Important to optain value from channel (to stop script).
    v := <-done
    fmt.Printf("%+v\n", v)
}

/*
working...done
false
*/
