package main

import (
  "fmt"
  "bytes"
  "encoding/json"
)

func main() {
  f1()
  f2()
}

func f1() {
  var b2 bytes.Buffer
  data := "x < y"

  enc := json.NewEncoder(&b2)
  enc.SetEscapeHTML(false)
  enc.Encode(data)
  
  fmt.Println("f1: ", b2.String())
}

func f2() {
  var data = []byte(`{"status": 200}`)
  var result map[string]interface{}

  decoder := json.NewDecoder(bytes.NewReader(data))
  decoder.UseNumber() // !!!

  err := decoder.Decode(&result)
  if err != nil {
    fmt.Println("error:", err)
    return
  }

  var status,_ = result["status"].(json.Number).Int64() //ok
  fmt.Println("f2 status value:",status)
}
