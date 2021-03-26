package main

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

type Item struct {
	Name  string
	Count int
	Type  string
}

type Data struct {
	Name      string
	CreatedAt int64
	Items     []Item
}

func main() {
	d := Data{
		Name:      "Example 1",
		CreatedAt: time.Now().Unix(),
		Items:     []Item{{Name: "ball", Type: "a", Count: 2}, {Name: "pump", Type: "b", Count: 1}},
	}

	//res := naive(d)
	res := withChan(d)
	fmt.Printf("\n%s\n", res)
}

var templateJSON = `{
	"name": "{{ .data.Name }}",
	"ts": {{ .data.CreatedAt }},
	"items": [
		{{ range $i, $item := .items }}
			{"name": "{{ $item.Name }}", "count":{{ $item.Count }}}
			{{ if not (eq $i $.n) }},{{ end }}
		{{ end }}
	]
}`

var tmpl = template.Must(template.New("tmpl").Parse(templateJSON))

func naive(data Data) string {
	var buf bytes.Buffer
	err := tmpl.Execute(
		&buf, map[string]interface{}{"data": data, "n": len(data.Items) - 1, "items": data.Items},
	)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

func withChan(data Data) string {
	ch := make(chan Item, 3)
	result := make(chan string)

	go produceItems(data, ch)
	go consumeItems(data, ch, result)

	return <-result
}

func produceItems(data Data, ch chan Item) {
	for _, item := range data.Items {
		ch <- item
	}
	close(ch)
}

func consumeItems(data Data, ch chan Item, result chan string) {
	var buf bytes.Buffer
	err := tmpl.Execute(
		&buf, map[string]interface{}{"data": data, "n": len(data.Items) - 1, "items": ch},
	)
	if err != nil {
		panic(err)
	}

	result <- buf.String()
}
