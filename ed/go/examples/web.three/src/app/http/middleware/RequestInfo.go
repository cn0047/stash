package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type RequestInfo struct {
	Next http.Handler
}

func (gm *RequestInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	fmt.Printf("%s %s %s \n", time.Now(), r.Method, r.URL.Path)
}
