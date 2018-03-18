package main

import (
    "encoding/json"
    "net/http"
    "strings"
)

type Response struct {
    Id string
    Msg string
}

func main() {
    http.HandleFunc("/v1/id/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != "GET" {
            w.Write([]byte(`{error:"501 Not Implemented"}`))
            return
        }

        url := r.URL.Path
        p := strings.LastIndex(url, "/") + 1
        id := url[p:]

        res := Response{Id: id, Msg: "OK."}
        e := json.NewEncoder(w)
        e.Encode(res)
    })
    http.ListenAndServe(":8080", nil)
}
