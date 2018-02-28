package middleware

import (
	"net/http"
)

type My struct {
	Next http.Handler
}

func (gm *My) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}
	println("My middleware.")
}
