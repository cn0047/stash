package experiment

import (
	"cloud.google.com/go/errorreporting"
	"cloud.google.com/go/logging"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"

	"go-app/config"
)

func StackDriverErrorsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	errorClient, err := errorreporting.NewClient(ctx, config.ProjectID, errorreporting.Config{
		ServiceName: "default",
		OnError: func(err error) {
			log.Errorf(ctx, "[20180703-001] Error client error: %v", err)
		},
	})
	if err != nil {
		log.Errorf(ctx, "[20180703-002] Filed to create new error client, error: %v", err)
	}

	defer errorClient.Close()
	defer errorClient.Flush()

	er := errors.New("My test error.")
	errorClient.Report(errorreporting.Entry{Error: er})

	fmt.Fprint(w, "Error reported.")
}

func StackDriverLogsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	client, err := logging.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Errorf(ctx, "[20180703-008] Filed to create new logging client, error: %v", err)
	}

	defer client.Close()

	logger := client.Logger("experiment").StandardLogger(logging.Info)
	logger.Println("My test log.")

	fmt.Fprint(w, "Log stored.")
}
