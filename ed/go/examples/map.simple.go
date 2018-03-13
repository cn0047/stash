package main

import "fmt"

func main() {
    lookup := make(map[string]int)
    lookup["goku"] = 9001
    power, exists := lookup["vegeta"]
    fmt.Println("1:", lookup, power, exists)

    m2 := map[string]string {
        "Bob": "Mr",
        "Amy": "Dr",
    }
    fmt.Println("2:", m2)
    delete(m2, "Amy")
    fmt.Println("3:", m2)
    fmt.Printf("4: %#v\n", m2["J"])
}
