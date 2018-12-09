package main

import (
  "bytes"
  "encoding/json"
  "fmt"
  "math/rand"
  "net/http"
  "time"
)

func random(min int, max int) int {
  rand.Seed(time.Now().UnixNano())
  return rand.Intn(max-min) + min
}

func main() {
  id := random(1, 7000)
  url := "https://realtimelog.herokuapp.com/ping"
  for {
    at := time.Now().UTC()
    j, _ := json.Marshal(map[string]interface{}{"id": id, "at": at})
    http.Post(url, "application/json", bytes.NewBuffer(j))
    fmt.Printf("Please open: %s to see new message, at: %s \n", url, at)
    time.Sleep(2 * time.Second)
  }
}
