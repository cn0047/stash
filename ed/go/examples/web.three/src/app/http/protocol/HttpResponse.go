package protocol

import (
	"encoding/json"
	"net/http"
)

// Canonical way to sends HTTP message response to client.
func HttpResponse(w http.ResponseWriter, response HttpMessage) {
	w.WriteHeader(getHttpCode(response))
	w.Header().Add("Content-Type", "application/json");

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		panic("RUNTIME-ERROR-JSON-1")
	}
}

// Gets HTTP status code from provided HttpMessage.
func getHttpCode(response HttpMessage) int {
	var httpCode int

	errorCode := response.GetError().Code
	successCode := response.GetSuccess().Code
	if successCode > 0 {
		httpCode = successCode
	}
	// This block was written in this way intentionally,
	// because error has higher priority and may overlap success code.
	if errorCode > 0 {
		httpCode = errorCode
	}

	return httpCode
}
