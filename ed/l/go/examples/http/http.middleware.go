package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error 3: %+v \n", err)
		}
		err = r.Body.Close()
		if err != nil {
			fmt.Printf("error 4: %+v \n", err)
		}
		r.Body = ioutil.NopCloser(bytes.NewReader(b))
		fmt.Printf("request body: %q\n", b)
		next.ServeHTTP(w, r)
	})
}

func middlewareH(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := ""
		for h, v := range r.Header {
			headers += fmt.Sprintf("%q: %q; ", h, v)
		}
		fmt.Printf("request headers: %s\n", headers)
		next.ServeHTTP(w, r)
	})
}

func middlewareT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("⏱ %s\n", time.Now())
		next.ServeHTTP(w, r)
	})
}

func middlewareR(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request method: %s, request URI: %s \n", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	rootHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	postHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("error 1: %+v \n", err)
		}
		err = r.Body.Close()
		if err != nil {
			fmt.Printf("error 2: %+v \n", err)
		}
		fmt.Printf("️request body: %s\n", b)
	})

	http.Handle("/", middlewareT(middlewareR(middlewareH(rootHandler))))
	http.Handle("/post", middlewareT(middlewareR(middlewareH(middlewareB(postHandler)))))
	http.ListenAndServe(":8080", nil)
}
