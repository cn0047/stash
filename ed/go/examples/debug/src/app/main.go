package main

import (
	"fmt"
	"net/http"
)

func main() {
	one()
}

func one() {
	web()
}

func web() {
	msg := "Hello world!"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	})
	http.ListenAndServe(":8080", nil)
}

func cli() {
	name, ending := "World", "!"
	ending2 := `)`
	fmt.Printf("Hello %s %s%s \n", name, ending, ending2)
}
