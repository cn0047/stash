package experiment

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/service/visit"
)

func DatastoreVisitHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	k, err := visit.TrackVisit(ctx, r)
	if err == nil {
		log.Infof(ctx, "✅ Visit already tracked, key: %v.", k)
	} else {
		log.Errorf(ctx, "❌ Failed to track visit, error: %v.", err)
	}

	visitsCount, err := visit.GetCount(ctx)
	if err == nil {
		fmt.Fprintf(w, "<br>This is visit # %v.", visitsCount)
	}
}
