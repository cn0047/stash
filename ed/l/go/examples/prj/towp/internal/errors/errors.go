package errors

import (
	"encoding/json"
	"errors"
	"github.com/to-com/wp/internal/business/validator"
	"net/http"
)

var (
	// ErrHTTPUnauthorized represents HTTP unauthorized client error.
	ErrHTTPUnauthorized = errors.New("HTTP error unauthorized")
)

type Error struct {
	Msg             string   `json:"message"`
	Cutoff          string   `json:"cutoff,omitempty"`
	ErrorFields     []string `json:"errorFields,omitempty"`
	ValidationError bool     `json:"validation_error,omitempty"`
}

type ResponseWithErrors struct {
	Errors []Error `json:"errors"`
}

func WriteError(w http.ResponseWriter, msg string, statusCode int) {
	errors := ResponseWithErrors{
		[]Error{
			{
				Msg: msg,
			},
		},
	}
	write(w, errors, statusCode)
}

func WritePlanValidationErrors(w http.ResponseWriter, err *validator.ValidationError, statusCode int) {
	var errs []Error
	for _, we := range err.WavesErrors {
		errs = append(errs, Error{
			Msg:             we.Error(),
			Cutoff:          we.Cutoff,
			ErrorFields:     we.ErrFields,
			ValidationError: true,
		})
	}
	write(w, ResponseWithErrors{Errors: errs}, statusCode)
}

func write(w http.ResponseWriter, responseWithErrors ResponseWithErrors, statusCode int) {
	encoded, err := json.Marshal(responseWithErrors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Unable to serialize response body"))

		return
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(encoded)
}
