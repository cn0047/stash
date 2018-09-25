package main

func main() {
    x := "foo"
    defer println(x)
    x = "bar"
}

// Result:
// foo
