package internal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/config"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPClient(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		t.Error(err)
	}

	httpClient := NewHTTPClient(cfg)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, server.URL, nil)
	if err != nil {
		t.Error(err)
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "ok", string(rawBody))
}
