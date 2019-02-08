package main

import (
  "encoding/json"
  "fmt"
  "net/http"
)

type Foo struct {
  Foo string `json:"foo"`
}

func main() {
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    decoder := json.NewDecoder(r.Body)
    f := Foo{}
    if err := decoder.Decode(&f); err != nil {
      fmt.Fprintf(w, "error: %s", err)
      return
    }
    fmt.Fprintf(w, "got foo: %s \n", f.Foo)
  })

  http.ListenAndServe(":8000", nil)
}

// curl -XPOST 'http://localhost:8000' -H 'Content-Type: application/json' -d '{"foo":"bar"}'
