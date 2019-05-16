package main

func main() {
	var m map[string]int
	m["x"] = 1 // panic: assignment to entry in nil map
}
