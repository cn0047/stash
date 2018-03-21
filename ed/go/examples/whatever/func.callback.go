package main

import "fmt"

func main() {
    foo(prn)
}

func foo(cb func(string)) {
    cb("This is from foo.")
}

func prn(m string) {
    fmt.Print(m)
}
