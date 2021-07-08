package main

func main() {
	// f1()
	// getValueFromNilMap(nil)
	getLenFromNilMap(nil)
}

func getLenFromNilMap(m map[string]int) {
	println(len(m)) // 0
}

func getValueFromNilMap(m map[string]int) {
	v, ok := m["key"]
	println(m == nil, v, ok) // true 0 false
}

func f1() {
	var m map[string]int
	m["x"] = 1 // panic: assignment to entry in nil map
}
