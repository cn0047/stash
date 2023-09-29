package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	p := ":8080"
	if len(os.Getenv("APP_PORT")) > 1 {
		p = ":" + os.Getenv("APP_PORT")
	}

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format("15:04:05") // time.TimeOnly
		fmt.Printf("req [%s]\n", t)
		w.Write([]byte(fmt.Sprintf("[%s] Hello world!\n", t)))
	})

	s := http.Server{
		Addr:         p,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}
	fmt.Printf("Serving on: http://localhost%s/\n", p)
	log.Fatal(s.ListenAndServe())
}
