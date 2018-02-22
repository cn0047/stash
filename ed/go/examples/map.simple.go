package main

import "fmt"

func main() {
    lookup := make(map[string]int)
    lookup["goku"] = 9001
    power, exists := lookup["vegeta"]
    fmt.Println(lookup, power, exists)

    m2 := map[string]string {
        "Bob": "Mr",
        "Amy": "Dr",
    }
    fmt.Println(m2)
    delete(m2, "Amy")
    fmt.Println(m2)
}
