package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheckHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/health-check", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HealthCheckHandler)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("ERROR-1 Got %v want %v", status, http.StatusOK)
    }

    expected := `{"alive": true}`
    if rr.Body.String() != expected {
        t.Errorf("ERROR-2 Got %v want %v", rr.Body.String(), expected)
    }
}
