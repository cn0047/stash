package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"ptt/app/payload"
	"ptt/app/ports"
	"ptt/config"
)

// Test_Handler represents test suite to test Handler.
func Test_Handler(ts *testing.T) {
	// Init test application.
	app, err := New(&config.Config{}, logrus.New())
	assert.Nil(ts, err)
	handler := &MainHTTPHandler{App: app}

	ts.Run("should return short ID error", func(tc *testing.T) {
		// Arrange.
		port := payload.Port{ID: "1"}
		data, err := json.Marshal(port)
		assert.Nil(tc, err)
		req, err := http.NewRequest("POST", "/v1/ports", bytes.NewBuffer(data))
		assert.Nil(tc, err)

		// Act.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CreatePortHandler)
		handler.ServeHTTP(rr, req)

		// Assert.
		assert.Contains(tc, rr.Body.String(), ports.ErrPortIDLength.Error())
	})

	ts.Run("success case", func(tc *testing.T) {
		// Arrange.
		data, err := json.Marshal(getDefaultPort())
		assert.Nil(tc, err)
		req, err := http.NewRequest("POST", "/v1/ports", bytes.NewBuffer(data))
		assert.Nil(tc, err)

		// Act.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.CreatePortHandler)
		handler.ServeHTTP(rr, req)

		// Assert.
		assert.Equal(tc, http.StatusOK, rr.Code)
	})
}

func getDefaultPort() payload.Port {
	return payload.Port{
		ID:          "UAODS",
		Name:        "Odessa",
		City:        "Odessa",
		Country:     "Ukraine",
		Alias:       []string{"ua_ods"},
		Regions:     []string{"Europe"},
		Coordinates: []float32{30.7233095, 46.482526},
		Province:    "Odessa Oblast",
		Timezone:    "Europe/Kiev",
		Unlocs:      []string{"UAODS"},
		Code:        46275,
	}
}
