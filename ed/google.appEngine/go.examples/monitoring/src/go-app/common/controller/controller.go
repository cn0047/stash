package controller

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
	"net/http"
)

// InternalResponse helper func which sends response for internal requests.
func InternalResponse(ctx context.Context, w http.ResponseWriter, action string, data interface{}, err error) {
	if err == nil {
		w.WriteHeader(http.StatusOK)
		log.Infof(ctx, "[✅] Performed action: %s, result: %v", action, data)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf(ctx, "[❌] Filed to perform action: %s, error: %v", action, err)
	}
}
