package main

import "fmt"

func main() {
    i := 4
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    case 4, 5:
        fmt.Println("4 of 5")
    default:
        fmt.Println("oops...")
    }
}
