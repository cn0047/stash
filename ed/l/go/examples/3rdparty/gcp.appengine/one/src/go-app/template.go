package go_app

import (
	"html/template"
	"net/http"
)

func templateHandler(w http.ResponseWriter, r *http.Request) {
	simple(w, r)
}

func simple(w http.ResponseWriter, r *http.Request) {
	dir := "template"

	tmpl := template.Must(template.ParseFiles(dir + "/main.html"))
	data := map[string]string{
		"title": "test",
	}
	tmpl.Execute(w, data)
}
