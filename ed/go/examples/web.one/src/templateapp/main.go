package main

import (
    "html/template"
    "net/http"
)

func main() {
    templates := populateTemplates()

    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        requestedFile := req.URL.Path[1:]
        template := templates.Lookup(requestedFile + ".html")
        if template != nil {
            template.Execute(res, nil)
        } else {
            res.WriteHeader(404)
        }
    })
    http.Handle("/img/", http.FileServer(http.Dir("public")))
    http.Handle("/css/", http.FileServer(http.Dir("public")))

    http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
    result := template.New("templates")
    const basePath = "templates"
    template.Must(result.ParseGlob(basePath + "/*.html"))
    return result
}
