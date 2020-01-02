package one

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(svc Service) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeHTTPError),
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		// r.Use(middleware)

		r.Method("GET", "/v1/hello/{key}", httptransport.NewServer(
			makeHelloEndpoint(svc),
			decodeHelloRequest,
			encodeHTTPResponse,
			options...))
	})

	return r
}

func decodeHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	param := chi.URLParam(r, "param")

	if param == "" {
		return "", errors.New("param must not be blank")
	}

	return param, nil
}

func encodeHTTPResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	return json.NewEncoder(w).Encode(map[string]interface{}{"message": response})
}

func encodeHTTPError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
}
