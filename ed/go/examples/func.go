package main

import "fmt"

func main() {
    r := foo("FOO")
    println(r)
    
    r2 := foo2("FOO2")
    println(r2)

    println(v(1, 2, 3))
}

func foo(m string) (string) {
    return "RESULT: " + m
}

func foo2(m string) (r string) {
    r = "RESULT 2: " + m
    return
}

func v(args ...int) (int) {
    fmt.Print(args)
    return len(args)
}
