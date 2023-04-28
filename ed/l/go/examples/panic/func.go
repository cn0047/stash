package main

func main() {
	var f func()
	f() // panic: runtime error: invalid memory address or nil pointer dereference
}
