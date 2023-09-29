package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	p := ":8081"
	if len(os.Getenv("APP_PORT")) > 1 {
		p = ":" + os.Getenv("APP_PORT")
	}
	targetURI := "http://localhost:8080"
	if len(os.Getenv("TARGET_URI")) > 1 {
		targetURI = os.Getenv("TARGET_URI")
	}

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Format(time.Kitchen)
		s, err := getData(targetURI)
		fmt.Printf("req [%s] err: %v\n", t, err)
		w.Write([]byte(fmt.Sprintf("got: %s", s)))
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

func getData(targetURI string) (string, error) {
	req, err := http.NewRequest("GET", targetURI, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new request, error: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to perform request, error: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body, error: %w", err)
	}

	return string(data), nil
}
