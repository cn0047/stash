package route

import (
	"net/http"

	"go-app/config"
	"go-app/controller/realtimelog"
)

func Realtimelog() {
	http.HandleFunc("/realtimelog/ping", realtimelog.PingHandler)
	http.HandleFunc("/realtimelog/pinging", realtimelog.PingingHandler)

	http.HandleFunc("/cronTask/realtimelog/ping", realtimelog.CronTaskPingHandler)
	http.HandleFunc("/cronTask/realtimelog/pinging", realtimelog.CronTaskPingingHandler)
	http.HandleFunc("/cronTask/realtimelog/addPingJob", realtimelog.CronTaskAddPingJobHandler)
	http.HandleFunc("/cronTask/realtimelog/addPingingJob", realtimelog.CronTaskAddPingingJobHandler)

	http.HandleFunc(config.WorkerPathPing, realtimelog.WorkerPingHandler)
	http.HandleFunc(config.WorkerPathPinging, realtimelog.WorkerPingingHandler)
}
