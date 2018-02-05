package main

import "fmt"

func main() {
    lookup := make(map[string]int)
    lookup["goku"] = 9001
    power, exists := lookup["vegeta"]

    fmt.Println(lookup, power, exists)
}
