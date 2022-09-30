package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"quizengine/app/payload"
)

// InitRoutes initializes all HTTP routes.
func InitRoutes(a *App, rootRouter *mux.Router) error {
	mainHandler := &MainHTTPHandler{
		App:       a,
		validator: validator.New(),
	}

	// Init root routes.
	rootRouter.HandleFunc("/health", mainHandler.HealthHandler).Methods("GET")
	rootRouter.HandleFunc("/auth", mainHandler.AuthenticateUserHandler).Methods("POST")
	rootRouter.HandleFunc("/users", mainHandler.CreateUserHandler).Methods("POST")

	// Init API routes with middlewares.
	r := rootRouter.PathPrefix("/v1").Subrouter()
	r.Use(authMiddleware)
	r.HandleFunc("/quizzes", mainHandler.ListQuizzesHandler).Methods("GET")
	r.HandleFunc("/quizzes/{id}", mainHandler.GetQuizHandler).Methods("GET")
	r.HandleFunc("/quizzes/{id}", mainHandler.UpsertQuizHandler).Methods("PUT")
	r.HandleFunc("/quizzes/{id}", mainHandler.DeleteQuizHandler).Methods("DELETE")
	r.HandleFunc("/submissions", mainHandler.ListSubmissionsHandler).Methods("GET")
	r.HandleFunc("/submissions", mainHandler.CreateSubmissionHandler).Methods("POST")

	return nil
}

// MainHTTPHandler holds all handlers for all HTTP routes.
type MainHTTPHandler struct {
	App       *App
	validator *validator.Validate
}

// HealthHandler represents main HTTP handler to check general application health.
func (h *MainHTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	sendSuccess(w, http.StatusOK, map[string]string{"status": "ok", "version": h.App.Config.BuildCommitHash})
}

// AuthenticateUserHandler represents main HTTP handler to authenticate user.
func (h *MainHTTPHandler) AuthenticateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to read body.")
		return
	}
	req := &payload.AuthenticateUserRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to unmarshal body.")
		return
	}
	err = h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	res, err := h.App.QuizEngine.AuthenticateUser(r.Context(), req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed create user, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusOK, res)
}

// CreateUserHandler represents main HTTP handler to create user.
func (h *MainHTTPHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to read body.")
		return
	}
	req := &payload.CreateUserRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to unmarshal body.")
		return
	}
	err = h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	err = h.App.QuizEngine.CreateUser(r.Context(), req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create user, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusCreated, payload.CreateUserResponse{})
}

// ListQuizzesHandler represents main HTTP handler to get quizzes.
func (h *MainHTTPHandler) ListQuizzesHandler(w http.ResponseWriter, r *http.Request) {
	req := &payload.ListQuizzesRequest{
		Author: strings.Join(r.URL.Query()["author"], ""),
	}
	err := h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	res, err := h.App.QuizEngine.ListQuizzes(r.Context(), req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to list quizzes, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusOK, res)
}

// GetQuizHandler represents main HTTP handler to get quiz.
func (h *MainHTTPHandler) GetQuizHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	res, err := h.App.QuizEngine.GetQuiz(r.Context(), id)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to get quiz, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusOK, res)
}

// UpsertQuizHandler represents main HTTP handler to upsert quiz.
func (h *MainHTTPHandler) UpsertQuizHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to read body.")
		return
	}
	req := &payload.UpsertQuizRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to unmarshal body.")
		return
	}
	req.ID = mux.Vars(r)["id"]
	req.Author = ctx.Value(payload.ContextKeyUserID).(string)
	err = h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	err = h.App.QuizEngine.UpsertQuiz(ctx, req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to update quiz, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusNoContent, payload.UpsertQuizResponse{})
}

// DeleteQuizHandler represents main HTTP handler to delete quiz.
func (h *MainHTTPHandler) DeleteQuizHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &payload.DeleteQuizRequest{
		Requester: ctx.Value(payload.ContextKeyUserID).(string),
		QuizID:    mux.Vars(r)["id"],
	}
	err := h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	err = h.App.QuizEngine.DeleteQuiz(r.Context(), req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to delete quiz, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusNoContent, payload.DeleteQuizResponse{})
}

// ListSubmissionsHandler represents main HTTP handler to get submissions.
func (h *MainHTTPHandler) ListSubmissionsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := &payload.ListSubmissionsRequest{
		Requester:  ctx.Value(payload.ContextKeyUserID).(string),
		UserID:     strings.Join(r.URL.Query()["user_id"], ""),
		QuizAuthor: strings.Join(r.URL.Query()["quiz_author"], ""),
	}
	err := h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	res, err := h.App.QuizEngine.ListSubmissions(ctx, req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to list submissions, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusOK, res)
}

// CreateSubmissionHandler represents main HTTP handler to create submissions.
func (h *MainHTTPHandler) CreateSubmissionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to read body.")
		return
	}
	req := &payload.CreateSubmissionRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Failed to unmarshal body.")
		return
	}
	req.UserID = ctx.Value(payload.ContextKeyUserID).(string)
	err = h.validator.Struct(req)
	if err != nil {
		sendError(w, http.StatusBadRequest, fmt.Sprintf("Failed validate request, err: %v", err))
		return
	}

	res, err := h.App.QuizEngine.CreateSubmission(ctx, req)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to create submission, err: %v", err))
		return
	}

	sendSuccess(w, http.StatusOK, res)
}

func sendSuccess(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)
	res, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("internal error: %v", err)
		res := fmt.Sprintf(`{"errors":[{"error":"%v","type":"internal error"}]}`, err)
		_, _ = w.Write([]byte(res))
		return
	}
	_, _ = w.Write(res)
}

func sendError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	t := "bad request"
	if code == http.StatusInternalServerError {
		t = "runtime error"
	}
	res := fmt.Sprintf(`{"errors":[{"error":"%s","type":"%s"}]}`, msg, t)
	_, _ = w.Write([]byte(res))
}
