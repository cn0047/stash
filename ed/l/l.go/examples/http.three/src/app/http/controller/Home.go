package controller

import (
	"net/http"

	"app/http/protocol"
)

type Home struct {
}

func (h Home) registerRoutes() {
	http.HandleFunc("/home", h.handleRequest)
}

func (h Home) handleRequest(w http.ResponseWriter, r *http.Request) {
	message := protocol.HttpSuccess(200, "This is REST-API for 'car' service.")
	protocol.HttpResponse(w, message)
}
