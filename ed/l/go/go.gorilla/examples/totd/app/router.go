package app

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"strings"

	"github.com/go-openapi/strfmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/to-com/poc-td/app/payload"
	"github.com/to-com/poc-td/app/swagger/restmodel"
	"github.com/to-com/poc-td/app/util"
)

var (
	//go:embed swagger/apidocs
	swaggerAssets embed.FS
)

// InitRoutes initializes all HTTP routes.
func InitRoutes(a *App, rootRouter *mux.Router) error {
	mainHandler := &MainHTTPHandler{App: a}

	// Swagger UI.
	swaggerAssetsContent, err := fs.Sub(fs.FS(swaggerAssets), "swagger/apidocs")
	if err != nil {
		return fmt.Errorf("failed to init FS for swagger assets, err: %w", err)
	}
	fileServer := http.FileServer(http.FS(swaggerAssetsContent))
	rootRouter.PathPrefix("/apidocs/").Handler(http.StripPrefix("/apidocs/", fileServer))

	// Init root routes.
	rootRouter.HandleFunc("/", mainHandler.HealthHandler).Methods("GET")
	rootRouter.HandleFunc("/health", mainHandler.HealthHandler).Methods("GET")

	// Init API routes with middlewares.
	r := rootRouter.PathPrefix("/v2").Subrouter()
	r.Use(headersMiddleware)
	r.Use(authMiddleware)
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/configs", mainHandler.GetConfigsHandler).Methods("GET")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/configs", mainHandler.UpdateConfigsHandler).Methods("PUT")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments", mainHandler.ListToteAssignmentsHandler).Methods("GET")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments", mainHandler.CreateToteAssignmentHandler).Methods("POST")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments", mainHandler.DeleteToteAssignmentHandler).Methods("DELETE")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments/{toteAssignmentId}", mainHandler.CreateToteAssignmentHandler).Methods("POST")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments", mainHandler.UpdateToteAssignmentHandler).Methods("PATCH")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments/{toteAssignmentId}", mainHandler.UpdateToteAssignmentHandler).Methods("PATCH")
	r.HandleFunc("/clients/{clientId}/mfcs/{mfcId}/tAssignments/{toteAssignmentId}", mainHandler.DeleteToteAssignmentHandler).Methods("DELETE")

	rootRouter.Handle("/", rootRouter)

	return nil
}

// MainHTTPHandler holds all handlers for all HTTP routes.
type MainHTTPHandler struct {
	App *App
}

// HealthHandler represents main HTTP handler to check general application health.
func (h *MainHTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.success200(w, map[string]string{"status": "ok", "version": h.App.Config.BuildCommitHash})
}

// GetConfigsHandler represents main HTTP handler to get MFCConfigs.
func (h *MainHTTPHandler) GetConfigsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := payload.GetConfigsInput{
		ClientID: mux.Vars(r)["clientId"],
		MfcID:    mux.Vars(r)["mfcId"],
	}
	out, err := h.App.MFCConfig.List(ctx, input)
	if err != nil {
		h.error500(w, "Failed to get configs.", err)
		return
	}

	h.success200(w, out)
}

// UpdateConfigsHandler represents main HTTP handler to update MFCConfigs.
func (h *MainHTTPHandler) UpdateConfigsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.error400(w, "Failed to read body.", "body")
		return
	}
	req := restmodel.UpdateConfigsRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.error400(w, "Failed to unmarshal body.", "body")
		return
	}
	if err := req.Validate(strfmt.Default); err != nil {
		h.error400(w, fmt.Sprintf("Failed to validate body, error: %v", err), "body")
		return
	}

	laneMapping, err := util.ConvertToMapWithInt64Keys[string, int64](req.Config.LaneMapping)
	if err != nil {
		h.error400(w, "Failed to convert laneMapping", "body")
	}
	expressLaneMapping, err := util.ConvertToMapWithInt64Keys[string, int64](req.Config.ExpressLaneMapping)
	if err != nil {
		h.error400(w, "Failed to convert expressLaneMapping", "body")
	}
	flowRacksMapping, err := util.ConvertToMapWithInt64Keys[string, string](req.Config.FlowRacksMapping)
	if err != nil {
		h.error400(w, "Failed to convert flowRacksMapping", "body")
	}

	input := payload.MFCConfig{
		ClientID: req.Config.ClientID,
		Env:      req.Config.Env,
		MfcID:    req.Config.MfcID,

		ErrorRamp:          req.Config.ErrorRamp,
		Count:              req.Config.Count,
		Depth:              req.Config.Depth,
		Start:              req.Config.Start,
		IDGen:              req.Config.IDGen,
		LaneMapping:        laneMapping,
		ExpressLaneMapping: expressLaneMapping,
		FlowRacksMapping:   flowRacksMapping,
	}
	if input.ClientID == "" {
		input.ClientID = mux.Vars(r)["clientId"]
	}
	if input.MfcID == "" {
		input.MfcID = mux.Vars(r)["mfcId"]
	}

	out, err := h.App.MFCConfig.Update(ctx, input)
	if err != nil {
		h.error500(w, "Failed to get configs.", err)
		return
	}

	h.success200(w, out)
}

// ListToteAssignmentsHandler represents main HTTP handler to list ToteAssignments.
func (h *MainHTTPHandler) ListToteAssignmentsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	input := payload.ListToteAssignmentsInput{
		ClientID: mux.Vars(r)["clientId"],
		MfcID:    mux.Vars(r)["mfcId"],
		View:     strings.Join(r.URL.Query()["view"], ""),
		OrderID:  strings.Join(r.URL.Query()["orderId"], ""),
	}
	out, err := h.App.ToteAssignment.List(ctx, input)
	if err != nil {
		h.error500(w, "Failed to list tote assignments.", err)
		return
	}

	h.success200(w, out)
}

// CreateToteAssignmentHandler represents main HTTP handler to create ToteAssignment.
func (h *MainHTTPHandler) CreateToteAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.error400(w, "Failed to read body.", "body")
		return
	}
	req := restmodel.CreateToteAssignmentRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.error400(w, "Failed to unmarshal body.", "body")
		return
	}
	if err := req.Validate(strfmt.Default); err != nil {
		h.error400(w, fmt.Sprintf("Failed to validate body, error: %v", err), "body")
		return
	}

	input := payload.CreateToteAssignmentInput{
		ClientID: mux.Vars(r)["clientId"],
		MfcID:    mux.Vars(r)["mfcId"],
		ID:       mux.Vars(r)["toteAssignmentId"],
		DryRun:   strings.Join(r.URL.Query()["dryRun"], "") == "true",
	}
	if req.ToteAssignment != nil {
		input.OrderID = req.ToteAssignment.OrderID
		input.ToteID = req.ToteAssignment.ToteID
		input.IsExpress = req.ToteAssignment.IsExpress
	}
	out, err := h.App.ToteAssignment.Create(ctx, input)
	if err != nil {
		h.error500(w, "Failed to create tote assignment.", err)
		return
	}

	h.success200(w, out)
}

// UpdateToteAssignmentHandler represents main HTTP handler to update ToteAssignment.
func (h *MainHTTPHandler) UpdateToteAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		h.error400(w, "Failed to read body.", "body")
		return
	}
	req := restmodel.UpdateToteAssignmentRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		h.error400(w, "Failed to unmarshal body.", "body")
		return
	}
	if err := req.Validate(strfmt.Default); err != nil {
		h.error400(w, fmt.Sprintf("Failed to validate body, error: %v", err), "body")
		return
	}

	input := payload.UpdateToteAssignmentInput{
		ClientID: mux.Vars(r)["clientId"],
		MfcID:    mux.Vars(r)["mfcId"],
		ID:       mux.Vars(r)["toteAssignmentId"],
	}
	if req.ToteAssignment != nil {
		input.LaneID = req.ToteAssignment.LaneID
		input.ToteID = req.ToteAssignment.ToteID
	}
	out, err := h.App.ToteAssignment.Update(ctx, input)
	if err != nil {
		h.error500(w, "Failed to update tote assignment.", err)
		return
	}

	h.success200(w, out)
}

func (h *MainHTTPHandler) DeleteToteAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	out, err := h.App.ToteAssignment.Delete(ctx, payload.DeleteToteAssignmentInput{
		ClientID: vars["clientId"],
		MfcID:    vars["mfcId"],
		ToteIDs:  r.URL.Query()["toteId"],
	})
	if err != nil {
		h.error500(w, "Failed to delete tote assignment.", err)
		return
	}
	h.success200(w, out)
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
	h.App.log.With(zap.Error(err), zap.String("msg", msg)).Errorf("%s, err: %v", msg, err)
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"","type":"internal error"}]}`, msg)
	_, _ = w.Write([]byte(res))
}
