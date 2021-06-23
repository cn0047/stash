package main

import (
	"fmt"
	"io"

	"github.com/valyala/fasttemplate"
)

func main() {
	// simpleExample()
	exampleWithFunc()
}

func simpleExample() {
	tpl := `x={{x}}`
	params := map[string]interface{}{
		"x": "007",
	}

	t := fasttemplate.New(tpl, "{{", "}}")
	s := t.ExecuteString(params)

	fmt.Printf("[simpleExample] %s \n", s)
}

func exampleWithFunc() {
	tpl := `x={{x}}`
	params := map[string]interface{}{
		"x": "007",
	}

	t, err := fasttemplate.NewTemplate(tpl, "{{", "}}")
	if err != nil {
		panic(err)
	}
	s := t.ExecuteFuncString(func(w io.Writer, tag string) (int, error) {
		v := params[tag].(string)
		return w.Write([]byte(v))
	})

	fmt.Printf("[exampleWithFunc] %s \n", s)
}
