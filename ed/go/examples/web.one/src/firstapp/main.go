package main

import (
	"errors"
	"net/http"

	"fmt"
	errorsWrapper "github.com/pkg/errors"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World from Go!"))
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusFound)
	})

	http.HandleFunc("/health-check", HealthCheckHandler)
	http.HandleFunc("/errors-wrap", ErrorsWrapHandler)

	http.ListenAndServe(":8000", nil)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"alive": true}`))
}

func ErrorsWrapHandler(w http.ResponseWriter, r *http.Request) {
	err := errors.New("my error")
	e := errorsWrapper.Wrap(err, "read failed")
	fmt.Printf("After wrap: ➡️ %+v ⬅️", e)
	w.Write([]byte("Look into console."))
	/*
		Result:

	After wrap: ➡️ my error
	read failed
	main.ErrorsWrapHandler
		/gh/ed/go/examples/web.one/src/firstapp/main.go:30
	net/http.HandlerFunc.ServeHTTP
		/usr/local/go/src/net/http/server.go:1947
	net/http.(*ServeMux).ServeHTTP
		/usr/local/go/src/net/http/server.go:2337
	net/http.serverHandler.ServeHTTP
		/usr/local/go/src/net/http/server.go:2694
	net/http.(*conn).serve
		/usr/local/go/src/net/http/server.go:1830
	runtime.goexit
		/usr/local/go/src/runtime/asm_amd64.s:2361 ⬅️
	*/
}
