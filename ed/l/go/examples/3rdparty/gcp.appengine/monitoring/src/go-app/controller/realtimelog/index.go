package realtimelog

import (
	"fmt"
	"google.golang.org/appengine"
	"net/http"

	"go-app/service/realtimelog"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	res, err := realtimelog.Ping(ctx, "ping")

	fmt.Fprintf(w, "Performed ping. <hr>Result: %v. <hr>Error: %v", res, err)
}

func PingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	res, err := realtimelog.Pinging(ctx, "pinging")

	fmt.Fprintf(w, "Performed pinging. <hr>Result: %v. <hr>Error: %v", res, err)
}
