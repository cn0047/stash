package route

import (
	"net/http"

	"go-app/controller/page"
)

func Page() {
	http.Handle("/page/", page.IndexHandler)
}
