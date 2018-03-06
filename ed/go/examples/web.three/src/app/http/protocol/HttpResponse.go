package protocol

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, response HttpMessage) {
	w.WriteHeader(getHttpCode(response))
	w.Header().Add("Content-Type", "application/json");

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	if err != nil {
		fmt.Printf("Error: %+v", err)
	}
}

func getHttpCode(response HttpMessage) int {
	var httpCode int

	errorCode := response.GetError().Code
	successCode := response.GetSuccess().Code
	if successCode > 0 {
		httpCode = successCode
	}
	if errorCode > 0 {
		httpCode = errorCode
	}

	return httpCode
}
