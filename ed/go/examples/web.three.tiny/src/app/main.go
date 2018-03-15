package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
)

type Response struct {
	Id string
}

func main() {
	http.HandleFunc("/v1/file-info/id/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		p := strings.LastIndex(url, "/") + 1
		id := url[p:]
		// or
		id2 := r.URL.Path[17:] // String "/v1/file-info/id/" length is 17
		fmt.Printf("Id = %#v\n", id2)

		res := Response{id}
		e := json.NewEncoder(w)
		e.Encode(res)
	})
	http.ListenAndServe(":8080", nil)
}
