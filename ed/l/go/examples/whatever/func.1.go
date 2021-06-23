package main

import (
    "fmt"
)

func main() {
    // println(foo("FOO"))   // RESULT: FOO
    // println(foo2("FOO2")) // RESULT 2: FOO2
    // println(v(1, 2, 3))   // [1 2 3]3
    x()
}

func x() {
    f1 := func(s string) {
        fmt.Printf("[f1] %s\n", s)
    }
    var f func(s string)
    f = f1
    addPrefix("test1", f)
    f = nil
    addPrefix("test2", f) // panic: runtime error: invalid memory address or nil pointer dereference
}

func addPrefix(s string, cb func(s string)) {
    str := "[with prefix] " + s
    cb(str)
}

func foo(m string) string {
    return "RESULT: " + m
}

func foo2(m string) (r string) {
    r = "RESULT 2: " + m
    return
}

func v(args ...int) int {
    fmt.Print(args)
    return len(args)
}
