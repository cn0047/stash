// This is 2 gen func.

package p

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("HelloWorldEntryPoint", HelloWorld)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("Hello World at: %s!", time.Now().Format(time.Kitchen))
	_, _ = fmt.Fprint(w, html.EscapeString(msg))
	log.Printf("%s", msg)
}
