package main

import (
	"html/template"
	"os"
)

func main() {
	dir := "/Users/k/web/kovpak/gh/ed/go/examples/template"

	tmpl := template.Must(template.ParseFiles(dir + "/main.html"))
	data := map[string]string{
		"title": "test",
	}
	tmpl.Execute(os.Stdout, data)
}
