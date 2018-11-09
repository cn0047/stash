package main

import "fmt"

type MyStruct struct {
	ID int
}

func main() {
	st := MyStruct{}
	fmt.Printf("Is empty: %#v \n", st == (MyStruct{}))
	st.ID = 1
	fmt.Printf("Is empty: %#v \n", st == (MyStruct{}))

  // Output:
  // Is empty: true
  // Is empty: false
}
