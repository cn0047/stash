package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	content := `
<html>
<head>
</head>
<body>
    Hello!
</body>
</html>
	`
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.Write([]byte(content))
	})
	p := ":8080"
	http.ListenAndServe(p, nil)
}
