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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	http.ListenAndServe(":8080", nil)
}

func cli() {
	name, ending := "World", "!"
	ending2 := `)`
	fmt.Printf("Hello %s %s%s \n", name, ending, ending2)
}
