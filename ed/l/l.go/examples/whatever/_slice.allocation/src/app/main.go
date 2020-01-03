package main

import (
	"app/lib"
	"fmt"
	"github.com/pkg/profile"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	n := 1000000
	//defer profile.Start(profile.MemProfile, profile.CPUProfile).Stop()
	defer profile.Start(profile.MemProfile).Stop()

	//cli(n)
	web(n)
}

func cli(n int) {
	//lib.F1(n)
	lib.F2(n)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("\n\n%+v\n\n", m)
}

func web(n int) {
	r := http.NewServeMux()

	r.HandleFunc("/f1", func(w http.ResponseWriter, r *http.Request) {
		lib.F1(n)
		w.Write([]byte("f1"))
	})

	r.HandleFunc("/f2", func(w http.ResponseWriter, r *http.Request) {
		lib.F2(n)
		w.Write([]byte("f2"))
	})

	// Register pprof handlers
	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8000", r)
}
