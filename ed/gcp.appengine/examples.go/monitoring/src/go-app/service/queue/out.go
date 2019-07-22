package queue

import (
	"golang.org/x/net/context"
	"net/http"

	"go-app/service/realtimelog"
)

func PerformPingJob(ctx context.Context, msg string) (r *http.Response, err error) {
	return realtimelog.Ping(ctx, msg)
}

func PerformPingingJob(ctx context.Context, msg string) (r map[int]*http.Response, err error) {
	return realtimelog.Pinging(ctx, msg)
}
