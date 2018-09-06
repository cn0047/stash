package main

import "fmt"

func main() {
  greeting, err := secureGreet("Dewey")
  if err != nil {
    fmt.Printf("error calling secureGreet: %v \n", err)
    return
  }
  fmt.Println(greeting)
}

func secureGreet(nephew string) (string, error) {
  var greeting string
  var err *NephewError
  if nephew != "Huey" && nephew != "Dewey" && nephew != "Louie" {
    greeting, err = "", &NephewError{nephew}
  } else {
    greeting, err = fmt.Sprint("Hello ", nephew), nil
  }
  return greeting, err
}

type NephewError struct {
  impostor string
}

func (e *NephewError) Error() string {
  return fmt.Sprint("I don't recall having a nephew named ", e.impostor)
}
