package main

import (
	"fmt"
	"html/template"
	"os"
)

var (
	dir = "/Users/k/web/kovpak/gh/ed/go/examples/template"
)

func main() {
	three()
}

func three() {
	t := template.New("three")
	funcs := template.FuncMap{
		"html": func(value interface{}) template.HTML {
			v := template.HTML(fmt.Sprint(value))
			return v
		},
	}
	templates := []string{
		dir + "/three.html",
		dir + "/threeAdd.html",
	}
	t, err := t.Funcs(funcs).ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"title": "Three",
		"txt":   `<table dir="ltr" border="1" cellspacing="0" cellpadding="0"><colgroup><col width="76"/><col width="90"/></colgroup><tbody><tr><td data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Active&#34;}">Active</td><td data-sheets-value="{&#34;1&#34;:2,&#34;2&#34;:&#34;Betsy&#34;}">Betsy</td></tr></tbody></table>`,
	}
	if err := t.ExecuteTemplate(os.Stdout, "tplThree", data); err != nil {
		panic(err)
	}
}

func two() {
	t := template.New("two")
	funcs := template.FuncMap{
		"plg": func(m string, data map[string]interface{}) (res string) {
			prf := data["type"].(string)
			return prf + m + "🎉"
		},
	}
	templates := []string{
		dir + "/one.html",
		dir + "/two.html",
		dir + "/twoAdd.html",
	}
	t, err := t.Funcs(funcs).ParseFiles(templates...)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"title":  "Two",
		"msg":    "It works!",
		"type":   "[OK] ",
		"public": true,
	}
	if err := t.ExecuteTemplate(os.Stdout, "tplTwo", data); err != nil {
		panic(err)
	}
}

func one() {
	tmpl := template.Must(template.ParseFiles(dir + "/one.html"))
	data := map[string]string{
		"title": "test",
	}
	tmpl.Execute(os.Stdout, data)
}
