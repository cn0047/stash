package main

func main() {
	m := map[string]int{}
	go func() {
		for {
			m["x"] = 1
		}
	}()
	for {
		// m["x"] = 1 // fatal error: concurrent map writes
		_ = m["x"] // fatal error: concurrent map read and map write
	}
}
