package main

func main() {
    f();
}

func f() {
    go func() {
        panic(500)
    }()
}
