package p

import (
	"fmt"
	"html"
	"net/http"
)

func MainFunc(w http.ResponseWriter, r *http.Request) {
	code := http.StatusOK
	msg := "ok"

	t := r.Header.Get("X-Token")
	if t == "" {
		code = http.StatusForbidden
		msg = "forbidden"
	}

	w.WriteHeader(code)
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
}
