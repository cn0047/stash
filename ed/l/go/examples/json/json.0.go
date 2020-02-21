package main

import (
  "bytes"
  "encoding/json"
  "fmt"
)

func main() {
  // f1()
  f2()
}

func f1() {
  var b2 bytes.Buffer
  data := "x < y"

  enc := json.NewEncoder(&b2)
  enc.SetEscapeHTML(false)
  enc.Encode(data)

  fmt.Println("f1: ", b2.String()) // f1:  "x < y"
}

func f2() {
  var data = []byte(`{"status": 200}`)
  decoder := json.NewDecoder(bytes.NewReader(data))
  decoder.UseNumber() // !!!

  var result map[string]interface{}
  err := decoder.Decode(&result)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  var status, _ = result["status"].(json.Number).Int64() //ok
  fmt.Println("f2 status value:", status)                // f2 status value: 200
}
