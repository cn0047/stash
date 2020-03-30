package main

func main() {
	switch x := 1; {
	case x < 2:
		println("a")
	case x < 3:
		println("b")
	default:
		println("z")
	}
}

// a
