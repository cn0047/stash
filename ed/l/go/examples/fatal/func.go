package main

func main() {
	var f func()
	go f() // fatal error: go of nil func value
}
