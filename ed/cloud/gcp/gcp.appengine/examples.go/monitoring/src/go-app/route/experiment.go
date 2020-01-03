package route

import (
	"net/http"

	"go-app/controller/experiment"
)

func Experiment() {
	http.HandleFunc("/experiment/datastore/visit", experiment.DatastoreVisitHandler)

	http.HandleFunc("/experiment/http/error500", experiment.HTTPError500Handler)

	http.HandleFunc("/experiment/runtime/numCPU", experiment.RuntimeNumCPUHandler)

	http.HandleFunc("/experiment/stackdriver/errors", experiment.StackDriverErrorsHandler)
	http.HandleFunc("/experiment/stackdriver/logs", experiment.StackDriverLogsHandler)
}
