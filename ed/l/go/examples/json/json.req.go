package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    // Read request body:
    // j, err := ioutil.ReadAll(r.Body)
    j, err := io.ReadAll(r.Body)
    if err != nil {
      fmt.Fprintf(w, "error: %s", err)
      return
    }
    defer r.Body.Close()

    fmt.Printf("Got JSON %s \n", j)
  })

  http.ListenAndServe(":8000", nil)
}

// curl -XPOST 'http://localhost:8000' -H 'Content-Type: application/json' -d '{"foo":"bar"}'
