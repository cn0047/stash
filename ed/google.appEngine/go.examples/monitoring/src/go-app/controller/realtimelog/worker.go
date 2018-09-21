package realtimelog

import (
	"google.golang.org/appengine"
	"net/http"

	"go-app/common/controller"
	"go-app/service/queue"
)

func WorkerPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	msg := r.FormValue("msg")
	res, err := queue.PerformPingJob(ctx, msg)

	controller.InternalResponse(ctx, w, "WorkerPing", res, err)
}

func WorkerPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	msg := r.FormValue("msg")
	res, err := queue.PerformPingingJob(ctx, msg)

	controller.InternalResponse(ctx, w, "WorkerPinging", res, err)
}
