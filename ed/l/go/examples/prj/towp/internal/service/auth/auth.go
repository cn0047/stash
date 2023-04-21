package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/common"
	internalErrors "github.com/to-com/wp/internal/errors"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Auth interface {
	CheckUser(context.Context, string) (string, error)
}

type Service struct {
	cfg        *config.Config
	logger     *zap.SugaredLogger
	httpClient *http.Client
}

func New(cfg *config.Config, logger *zap.SugaredLogger, httpClient *http.Client) *Service {
	return &Service{
		cfg:        cfg,
		logger:     logger,
		httpClient: httpClient,
	}
}

var _ Auth = &Service{}

type GetUserIDResponse struct {
	UserID       string `json:"user_id,omitempty"`
	ErrorMessage string `json:"error,omitempty"`
}

func (s *Service) CheckUser(ctx context.Context, retailerID string) (string, error) {
	ctx, span := trace.StartSpan(ctx, "auth.get-user-id")
	defer span.End()

	url := fmt.Sprintf(s.cfg.AuthURLTemplate, retailerID, common.GetCtxEnv(ctx)) +
		"/api/v1/user/permissions"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Add("X-Token", common.GetCtxToken(ctx))

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	statusOK := resp.StatusCode >= 200 && resp.StatusCode < 300
	statusUnauthorized := resp.StatusCode == http.StatusUnauthorized
	rawBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if statusOK || statusUnauthorized {
		var user GetUserIDResponse
		err := json.Unmarshal(rawBody, &user)
		if err != nil {
			return "", err
		}
		if statusUnauthorized {
			return "", fmt.Errorf("%w: %s", internalErrors.ErrHTTPUnauthorized, user.ErrorMessage)
		}

		return user.UserID, nil
	}

	return "", fmt.Errorf(
		"an error occurred while retrieving userId, response: %s", string(rawBody))
}
