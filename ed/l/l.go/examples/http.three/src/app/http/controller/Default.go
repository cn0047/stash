package controller

import (
	"net/http"

	"app/http/protocol"
)

type Default struct {
}

func (d Default) registerRoutes() {
	http.HandleFunc("/", d.handleRequest)
}

func (d Default) handleRequest(w http.ResponseWriter, r *http.Request) {
	message := protocol.HttpError(501, "Not Implemented.")
	protocol.HttpResponse(w, message)
}
