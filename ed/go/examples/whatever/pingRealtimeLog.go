package main

import (
  "bytes"
  "encoding/json"
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
  for {
    j, _ := json.Marshal(map[string]interface{}{"id": id, "at": time.Now().UTC()})
    http.Post("https://realtimelog.herokuapp.com/ping", "application/json", bytes.NewBuffer(j))
    time.Sleep(2 * time.Second)
  }
}
