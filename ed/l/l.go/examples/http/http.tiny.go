package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/v1/id/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		pos := strings.LastIndex(path, "/") + 1
		id := path[pos:]
		e := json.NewEncoder(w)
		e.Encode(map[string]string{"id": id})
	})

	http.ListenAndServe(":8080", nil)
}
