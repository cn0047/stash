package errors

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/internal/business/validator"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	respWithErrors = ResponseWithErrors{
		Errors: []Error{
			{
				Msg:             "test error msg, no cutoff",
				ValidationError: true,
			},
			{
				Msg:    "test error msg, cutoff specified",
				Cutoff: "01:00",
				ErrorFields: []string{
					"cutoff",
				},
				ValidationError: true,
			},
			{
				Msg:    "test error msg, fields specified",
				Cutoff: "02:00",
				ErrorFields: []string{
					"from",
					"to",
				},
				ValidationError: true,
			},
		},
	}
)

func TestWrite(t *testing.T) {
	httpResponse := httptest.NewRecorder()
	write(httpResponse, respWithErrors, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

	var writtenErrors ResponseWithErrors
	err := json.Unmarshal(httpResponse.Body.Bytes(), &writtenErrors)
	assert.Empty(t, err)
	assert.Equal(t, respWithErrors, writtenErrors)
}

func TestWriteError(t *testing.T) {
	testErrorMsg := "test error msg"
	httpResponse := httptest.NewRecorder()
	WriteError(httpResponse, testErrorMsg, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

	var writtenErrors ResponseWithErrors
	err := json.Unmarshal(httpResponse.Body.Bytes(), &writtenErrors)
	assert.Empty(t, err)
	assert.Equal(t, ResponseWithErrors{
		Errors: []Error{
			{
				Msg: testErrorMsg,
			},
		},
	}, writtenErrors)
}

func TestWritePlanValidationErrors(t *testing.T) {
	validationErr := validator.ValidationError{
		Err: fmt.Errorf("test validation error"),
		WavesErrors: []*validator.WaveError{
			{
				Err: fmt.Errorf("test error msg, no cutoff"),
			},
			{
				Err:    fmt.Errorf("test error msg, cutoff specified"),
				Cutoff: "01:00",
				ErrFields: []string{
					"cutoff",
				},
			},
			{
				Err:    fmt.Errorf("test error msg, fields specified"),
				Cutoff: "02:00",
				ErrFields: []string{
					"from",
					"to",
				},
			},
		},
	}

	httpResponse := httptest.NewRecorder()
	WritePlanValidationErrors(httpResponse, &validationErr, http.StatusBadRequest)

	assert.Equal(t, http.StatusBadRequest, httpResponse.Code)

	var writtenErrors ResponseWithErrors
	err := json.Unmarshal(httpResponse.Body.Bytes(), &writtenErrors)
	assert.Empty(t, err)
	assert.Equal(t, respWithErrors, writtenErrors)
}
