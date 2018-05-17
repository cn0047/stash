package main

import (
	"app/lib"
	"fmt"
	"github.com/pkg/profile"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	n := 1000000
	defer profile.Start(profile.MemProfile, profile.CPUProfile).Stop()

	cli(n)
}

func cli(n int) {
	//lib.F1(n)
	lib.F2(n)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\n\n%+v\n\n", m)
}

func web(n int) {
	http.HandleFunc("/f1", func(w http.ResponseWriter, r *http.Request) {
		lib.F1(n)
		w.Write([]byte("f1"))
	})

	http.HandleFunc("/f2", func(w http.ResponseWriter, r *http.Request) {
		lib.F2(n)
		w.Write([]byte("f2"))
	})

	http.ListenAndServe(":8080", nil)
}
