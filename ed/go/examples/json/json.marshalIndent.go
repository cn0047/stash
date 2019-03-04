package main

import (
  "encoding/json"
  "fmt"
)

type Entry struct {
    Key string `json:"key"`
}

func main() {
  e := Entry{Key: "value"}
  res, _ := json.Marshal(e)
  fmt.Println(string(res))

  res, _ = json.MarshalIndent(e, "", " - ")
  fmt.Println(string(res))
}

/*

{"key":"value"}
{
 - "key": "value"
}

*/
