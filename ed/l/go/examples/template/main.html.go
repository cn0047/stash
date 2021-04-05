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
	// getSliceElementByIndex()
	// orSimple()
	// three()
	// loop()
	// ifWithNotDefinedVar()
	// eqSimple()
	// orWithEq()
	// andSimple()
	// notSimple()
	// notWithAnd()
	// objAsParams()
	// getValueFromMapInMap()
	// invalidTemplate()
	// invalidTemplate()
	invalidTemplateParams()
}

func invalidTemplateParams() {
	t := `<br>[invalidTemplateParams] {{range $i,$u := .us}} {{$u.Name}} {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"us": "ok",
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		fmt.Printf("[invalidTemplateParams] execute error: %+v \n", err)
	}
}

func invalidTemplate() {
	t := "<br>[invalidTemplate]: {{ range $i, $unit := .units }}"
	_, err := template.New("tmpl").Parse(t)
	fmt.Printf("[invalidTemplate] parse error: %+v \n", err)
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

type params struct {
	Msg string
	Ok  bool
}

func objAsParams() {
	t := `<br>[objAsParams]: {{ if .Ok }} {{ .Msg }} {{ end }} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := params{Msg: "test", Ok: true}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func one() {
	tmpl := template.Must(template.ParseFiles(dir + "/one.html"))
	data := map[string]string{
		"title": "test",
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func getSliceElementByIndex() {
	t := `<br> Title: {{index .titles 0}}<hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"titles": []string{"test1", "test2"},
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func getValueFromMapInMap() {
	t := `<br>[getValueFromMapInMap] {{ .params.foo }}<hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"params": map[string]string{"foo": "bar"},
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func eqSimple() {
	t := `<br>[eqSimple] {{if eq .a 1}} ok {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": 1,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func orWithEq() {
	t := `<br>[orWithEq] {{if or (eq .a 1) (eq .b 2)}} TRUE {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": 1,
		"b": 2,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func orSimple() {
	t := `<br>[orSimple] {{if or .a .b}} TRUE {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": false,
		"b": true,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func andSimple() {
	t := `<br>[andSimple] {{if and .a .b}} TRUE {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": true,
		"b": true,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func notSimple() {
	t := `<br>[notSimple] {{if not .a}} TRUE {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": false,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func notWithAnd() {
	t := `<br>[notWithAnd] {{if and .a (not .b)}} TRUE {{end}} <hr>`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"a": true,
		"b": false,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func loop() {
	t := `List: {{range $k, $v := .list}} {{$k}}={{$v}}; {{end}}`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{
		"list": []string{"a", "b", "c"},
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}

func ifWithNotDefinedVar() {
	t := `ok: {{if .ok -}} v {{end}}`
	tmpl := template.Must(template.New("tmpl").Parse(t))
	data := map[string]interface{}{"no": "no"}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		panic(err)
	}
}
