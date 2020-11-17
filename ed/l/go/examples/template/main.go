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
	// f1()
	// f2()
	// three()
	// loop()
	ifWithNotDefinedVar()
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

func f1() {
	t := `<br> Title: {{index .titles 0}}<hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"titles": []string{"test1", "test2"},
	}
	tmpl.Execute(os.Stdout, data)
}

func f2() {
	t := `<br> {{if or .a .b}}[f2] TRUE{{end}}<hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": false,
		"b": true,
	}
	tmpl.Execute(os.Stdout, data)
}

func loop() {
	t := `List: {{range $k, $v := .list}} {{$k}}={{$v}}; {{end}}`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"list": []string{"a", "b", "c"},
	}
	tmpl.Execute(os.Stdout, data)
}

func ifWithNotDefinedVar() {
	t := `ok: {{if .ok -}} v {{end}}`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{"no": "no"}
	tmpl.Execute(os.Stdout, data)
}
