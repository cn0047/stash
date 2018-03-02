package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Response struct {
	Id string
}

func main() {
	http.HandleFunc("/v1/file-info/id/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path
		p := strings.LastIndex(url, "/") + 1
		id := url[p:]

		res := Response{id}
		e := json.NewEncoder(w)
		e.Encode(res)
	})
	http.ListenAndServe(":8080", nil)
}
