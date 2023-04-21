package internal

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/mocks"
	"go.uber.org/zap"
	"gopkg.in/h2non/gock.v1"
	"testing"
)

// TestApplication represents test application tailored for unit/integration tests.
type TestApplication struct {
	ctrl         *gomock.Controller
	mockBusiness *mocks.MockBusiness
	mockAuth     *mocks.MockAuthentication

	App *Application
}

// Shutdown stops/closes opened handlers for test application.
func (t *TestApplication) Shutdown() {
	t.ctrl.Finish()
	gock.Off()
}

func TestAppOk(t *testing.T) {
	logger := foundation.NewLogger()
	app, err := New(logger)

	assert.Nil(t, err)

	assert.IsType(t, &config.Config{}, app.cfg)
	assert.IsType(t, &zap.SugaredLogger{}, app.logger)

	assert.IsType(t, &HTTPHandler{}, app.handler)
}
