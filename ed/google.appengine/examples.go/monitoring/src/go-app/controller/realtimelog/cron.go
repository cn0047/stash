package realtimelog

import (
	"google.golang.org/appengine"
	"net/http"

	"go-app/common/controller"
	"go-app/service/queue"
	"go-app/service/realtimelog"
)

func CronTaskPingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	res, err := realtimelog.Ping(ctx, "ping-from-cron")

	controller.InternalResponse(ctx, w, "CronTaskPing", res, err)
}

func CronTaskPingingHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	res, err := realtimelog.Pinging(ctx, "pinging-from-cron")

	controller.InternalResponse(ctx, w, "CronTaskPinging", res, err)
}

func CronTaskAddPingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	err := queue.AddPingJob(ctx, "ping-from-queue")

	controller.InternalResponse(ctx, w, "CronTaskAddPingJob", "ok", err)
}

func CronTaskAddPingingJobHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	err := queue.AddPingingJob(ctx, "pinging-from-queue")

	controller.InternalResponse(ctx, w, "CronTaskAddPingingJob", "ok", err)
}
