//go:generate mockgen -source handler.go -destination mocks/mocks_for_handler.go -package mocks
package internal

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/to-com/wp/internal/business/validator"
	"go.uber.org/zap"
	"net/http"

	"github.com/to-com/go-telemetry/sdpropagation"
	"github.com/to-com/wp/internal/dto"
	ie "github.com/to-com/wp/internal/errors"
)

const (
	ContentTypeHeader = "Content-Type"
	JSONContentType   = "application/json"
)

type Business interface {
	Createwp(ctx context.Context, wp dto.wpRequest) (dto.wpResponse, error)
	Getwp(ctx context.Context, retailerID, mfcID string) (dto.wpResponse, error)
	GenerateTriggers(ctx context.Context) (res dto.GenerateTriggersResponse, err error)
	GetTriggers(ctx context.Context, retailerID, mfcID string) (dto.GetTriggersResponse, error)
	FireTriggers(ctx context.Context) (dto.FireTriggersResponse, error)
}

type Authentication interface {
	CheckUser(ctx context.Context, retailerID string) (string, error)
}

type HTTPHandler struct {
	logger   *zap.SugaredLogger
	Business Business
	Auth     Authentication
}

func NewHTTPHandler(logger *zap.SugaredLogger, bs Business, auth Authentication) *HTTPHandler {
	return &HTTPHandler{
		logger:   logger,
		Business: bs,
		Auth:     auth,
	}
}

func (h *HTTPHandler) writeError(w http.ResponseWriter, msg string, statusCode int) {
	h.logger.Errorf("HTTP handler error: %s, statusCode: %d", msg, statusCode)
	ie.WriteError(w, msg, statusCode)
}

func (h *HTTPHandler) Getwp(w http.ResponseWriter, r *http.Request) {
	ctx, span := sdpropagation.StartSpanWithRemoteParentFromRequest(r, "wp.get-wave-plan")
	defer span.End()

	ctx, err := PackInCtx(ctx, r, WithPackEnvInCtx)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	params := mux.Vars(r)
	retailerID, mfcID := params["retailerId"], params["mfcId"]

	wp, err := h.Business.Getwp(ctx, retailerID, mfcID)

	if err == nil && wp.ID == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	h.writeResponse(w, http.StatusOK, wp, err)
}

func (h *HTTPHandler) Createwp(w http.ResponseWriter, r *http.Request) {
	ctx, span := sdpropagation.StartSpanWithRemoteParentFromRequest(r, "wp.create-wave-plan")
	defer span.End()

	ctx, err := PackInCtx(ctx, r, WithPackTokenInCtx, WithPackEnvInCtx)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	params := mux.Vars(r)
	retailerID, mfcID := params["retailerId"], params["mfcId"]

	userID, authErr := h.Auth.CheckUser(ctx, retailerID)
	if authErr != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(authErr, ie.ErrHTTPUnauthorized) {
			statusCode = http.StatusUnauthorized
		}
		h.writeError(w, authErr.Error(), statusCode)

		return
	}

	var wp dto.wpRequest
	err = json.NewDecoder(r.Body).Decode(&wp)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	wp.RetailerID = retailerID
	wp.MfcID = mfcID
	wp.UserID = userID

	wpResponse, err := h.Business.Createwp(ctx, wp)
	h.writeResponse(w, http.StatusCreated, wpResponse, err)
}

func (h *HTTPHandler) writeResponse(w http.ResponseWriter, statusCode int, data any, err error) {
	w.Header().Set(ContentTypeHeader, JSONContentType)
	if err != nil {
		var (
			validationErr *validator.ValidationError
		)

		switch {
		case errors.As(err, &validationErr):
			ie.WritePlanValidationErrors(w, validationErr, http.StatusBadRequest)
			return
		default:
			h.writeError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	encoded, err := json.Marshal(data)
	if err != nil {
		h.writeError(w, "failed to marshal response data", http.StatusInternalServerError)

		return
	}
	w.WriteHeader(statusCode)
	_, _ = w.Write(encoded)
}

func (h *HTTPHandler) generateTriggers(w http.ResponseWriter, r *http.Request) {
	ctx, span := sdpropagation.StartSpanWithRemoteParentFromRequest(r, "wp.generate-triggers")
	defer span.End()

	ctx, err := PackInCtx(ctx, r, WithPackEnvInCtx)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	res, err := h.Business.GenerateTriggers(ctx)
	h.writeResponse(w, http.StatusOK, res, err)
}

func (h *HTTPHandler) GetTriggers(w http.ResponseWriter, r *http.Request) {
	ctx, span := sdpropagation.StartSpanWithRemoteParentFromRequest(r, "wp.get-triggers")
	defer span.End()

	ctx, err := PackInCtx(ctx, r, WithPackEnvInCtx)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	retailerID := r.URL.Query().Get("retailerId")
	mfcID := r.URL.Query().Get("mfcId")

	triggers, err := h.Business.GetTriggers(ctx, retailerID, mfcID)
	h.writeResponse(w, http.StatusOK, triggers, err)
}

func (h *HTTPHandler) FireTriggers(w http.ResponseWriter, r *http.Request) {
	ctx, span := sdpropagation.StartSpanWithRemoteParentFromRequest(r, "wp.fire-triggers")
	defer span.End()

	ctx, err := PackInCtx(ctx, r, WithPackEnvInCtx)
	if err != nil {
		h.writeError(w, err.Error(), http.StatusBadRequest)

		return
	}

	triggers, err := h.Business.FireTriggers(ctx)
	h.writeResponse(w, http.StatusOK, triggers, err)
}
