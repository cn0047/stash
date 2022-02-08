package p

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func MainFunc(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hello World! Now: %s.", time.Now().Format(time.Kitchen))
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
	log.Printf("%s", msg)
}
