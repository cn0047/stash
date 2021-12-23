package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func success200(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		error500(w, "Failed to marshal response data.", err)
		return
	}

	_, _ = w.Write(res)
}

func error400(w http.ResponseWriter, msg string, ref string) {
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"%s","type":"bad request"}]}`, msg, ref)
	_, _ = w.Write([]byte(res))
}

func error500(w http.ResponseWriter, msg string, err error) {
	log.Printf("%s, err: %v", msg, err)
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"","type":"internal error"}]}`, msg)
	_, _ = w.Write([]byte(res))
}
