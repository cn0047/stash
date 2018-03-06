package middleware

import (
	"fmt"
	"net/http"
)

type OsInfo struct {
	Next http.Handler
}

func (gm *OsInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if gm.Next == nil {
		gm.Next = http.DefaultServeMux
	}

	fmt.Printf("%s \n", "OS info not specified.")
}
