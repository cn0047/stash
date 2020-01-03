package main

import (
	"fmt"
	"net/http"
)

type MyStruct struct {
	ID int
}

func main() {
	ok()
}

func ok() {
	st := MyStruct{}
	fmt.Printf("Is empty: %#v \n", st == (MyStruct{}))
	st.ID = 1
	fmt.Printf("Is empty: %#v \n", st == (MyStruct{}))

	// Output:
	// Is empty: true
	// Is empty: false
}

func err() {
	c := http.Client{}
	fmt.Println(c == (http.Client{})) // invalid operation: c == http.Client literal (struct containing func(*http.Request, []*http.Request) error cannot be compared)
}
