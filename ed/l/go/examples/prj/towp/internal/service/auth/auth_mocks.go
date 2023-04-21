package auth

import (
	"context"
	"fmt"
	"github.com/to-com/wp/internal/common"
	"gopkg.in/h2non/gock.v1"
	"net/http"
)

var (
	asBaseURLFormat = "https://as-%s-%s.tom.to.com"
	asPath          = "/api/v1/user/permissions"
)

const (
	UserID    = "OSdP9Zt7yeNguXDXjkkRjV7nL2b2"
	errorResp = `{"error": "failed to connect to auth"}`
	AuthErr   = "Invalid token string"
)

func MockCheckUserSuccessResponse(ctx context.Context, retailerID string) {
	asBaseURL := fmt.Sprintf(asBaseURLFormat, retailerID, common.GetCtxEnv(ctx))

	gock.New(asBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		MatchHeader("X-Token", common.GetCtxToken(ctx)).
		Get(asPath).
		Reply(http.StatusOK).
		File("testdata/get_user_permissions_response.json")
}

func MockCheckUserUnauthorizedResponse(ctx context.Context, retailerID string) {
	asBaseURL := fmt.Sprintf(asBaseURLFormat, retailerID, common.GetCtxEnv(ctx))

	gock.New(asBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(asPath).
		Reply(http.StatusUnauthorized).
		JSON(fmt.Sprintf(`{"error": "%s"}`, AuthErr))
}

func MockCheckUserErrorResponse(ctx context.Context, retailerID string) {
	asBaseURL := fmt.Sprintf(asBaseURLFormat, retailerID, common.GetCtxEnv(ctx))

	gock.New(asBaseURL).
		WithOptions(gock.Options{DisableRegexpHost: true}).
		Persist().
		Get(asPath).
		Reply(http.StatusInternalServerError).
		JSON(errorResp)
}
