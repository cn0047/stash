package main

func main() {
	// simple() // a
	// stringFail() // nothing printed
	stringOk() // a|b|c
}

func stringOk() {
	x := "a"

	switch x {
	case "a", "b", "c":
		println("a|b|c")
	case "d", "e":
		println("d|e")
	default:
		println("unknown")
	}
}

func stringFail() {
	x := "a"

	switch x {
	case "a":
	case "b":
	case "c":
		println("a|b|c")
	case "d":
	case "e":
		println("d|e")
	default:
		println("unknown")
	}
}

func simple() {
	switch x := 1; {
	case x < 2:
		println("a")
	case x < 3:
		println("b")
	default:
		println("z")
	}
}
