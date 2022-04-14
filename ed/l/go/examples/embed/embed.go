// @see: open http://localhost:3000
// @see: curl http://localhost:3000/health

package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
)

var (
	//go:embed desc.txt
	desc string
	//go:embed assets
	assets embed.FS
)

func main() {
	fmt.Printf("desc: %s", desc)

	web()
}

func web() {
	var staticFS = fs.FS(assets)
	htmlContent, err := fs.Sub(staticFS, "assets")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServer(http.FS(htmlContent))
	http.Handle("/", fileServer)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("It works!\n"))
	})

	port := "3000"
	log.Printf("Listening on :%s...\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
