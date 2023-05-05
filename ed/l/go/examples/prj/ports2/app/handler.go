package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"ptt/app/payload"
)

// MainHTTPHandler holds all handlers for all HTTP routes.
type MainHTTPHandler struct {
	App *App
}

// HealthHandler represents main HTTP handler to check general application health.
func (h *MainHTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.success200(w, map[string]string{"status": "ok"})
}

// CreatePortHandler represents HTTP handler to create port.
func (h *MainHTTPHandler) CreatePortHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.error400(w, "Failed to read body.", "body")
		return
	}
	req := &payload.Port{}
	err = json.Unmarshal(body, req)
	if err != nil {
		h.error400(w, "Failed to unmarshal body.", "body")
		return
	}
	if err := h.App.Validator.Struct(req); err != nil {
		h.error400(w, fmt.Sprintf("Failed to validate body, error: %v", err), "body")
		return
	}

	res, err := h.App.PortsService.CreatePort(ctx, req)
	if err != nil {
		h.error400(w, "Failed to create port. "+err.Error(), "body")
		return
	}

	h.success200(w, res)
}

// UpdatePortHandler represents HTTP handler to update port.
func (h *MainHTTPHandler) UpdatePortHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.error400(w, "Failed to read body.", "body")
		return
	}
	req := &payload.Port{}
	err = json.Unmarshal(body, req)
	if err != nil {
		h.error400(w, "Failed to unmarshal body.", "body")
		return
	}
	// Get port ID from URL path parameter.
	// @TODO/TBD: This logic should be mentioned in swagger.yaml.
	req.ID = mux.Vars(r)["portId"]
	if err := h.App.Validator.Struct(req); err != nil {
		h.error400(w, fmt.Sprintf("Failed to validate body, error: %v", err), "body")
		return
	}

	res, err := h.App.PortsService.UpdatePort(ctx, req)
	if err != nil {
		h.error400(w, "Failed to create port."+err.Error(), "body")
		return
	}

	h.success200(w, res)
}

func (h *MainHTTPHandler) success200(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.Marshal(data)
	if err != nil {
		h.error500(w, "Failed to marshal response data.", err)
		return
	}

	_, _ = w.Write(res)
}

func (h *MainHTTPHandler) error400(w http.ResponseWriter, msg string, ref string) {
	w.Header().Set("Content-Type", "application/json")
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"%s","type":"bad request"}]}`, msg, ref)
	_, _ = w.Write([]byte(res))
}

func (h *MainHTTPHandler) error500(w http.ResponseWriter, msg string, err error) {
	w.Header().Set("Content-Type", "application/json")
	h.App.Log.Errorf("%s, err: %v", msg, err)
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"","type":"internal error"}]}`, msg)
	_, _ = w.Write([]byte(res))
}
