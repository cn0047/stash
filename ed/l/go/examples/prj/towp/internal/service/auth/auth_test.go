package auth

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/common"
	internalErrors "github.com/to-com/wp/internal/errors"
	"gopkg.in/h2non/gock.v1"
	"net/http"
	"testing"
)

func prepareAuth(t *testing.T) *Service {
	t.Helper()

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unable to load config for testing Auth service, error: %v", err)
	}
	logger := foundation.NewLogger()
	httpClient := &http.Client{}

	return New(cfg, logger, httpClient)
}

func TestCheckUser(t *testing.T) {
	authS := prepareAuth(t)

	testRetailer := "maf"
	testEnv := "dev"

	t.Run("auth service returns user permissions", func(t *testing.T) {
		defer gock.Off()
		ctx := context.Background()
		ctx = context.WithValue(ctx, common.CtxKeyEnv, testEnv)

		MockCheckUserSuccessResponse(ctx, testRetailer)

		info, err := authS.CheckUser(ctx, testRetailer)

		assert.Empty(t, err)
		assert.Equal(t, UserID, info)
	})

	t.Run("auth service returns 401", func(t *testing.T) {
		defer gock.Off()
		ctx := context.Background()
		ctx = context.WithValue(ctx, common.CtxKeyEnv, testEnv)

		MockCheckUserUnauthorizedResponse(ctx, testRetailer)

		info, err := authS.CheckUser(ctx, testRetailer)

		assert.Empty(t, info)
		assert.Equal(t, fmt.Sprintf("%v: %s", internalErrors.ErrHTTPUnauthorized, AuthErr),
			err.Error())
	})

	t.Run("auth service returns other error", func(t *testing.T) {
		defer gock.Off()
		ctx := context.Background()
		ctx = context.WithValue(ctx, common.CtxKeyEnv, testEnv)

		MockCheckUserErrorResponse(ctx, testRetailer)

		info, err := authS.CheckUser(ctx, testRetailer)

		assert.Empty(t, info)
		assert.Equal(t, "an error occurred while retrieving userId, response: "+errorResp, err.Error())
	})
}
