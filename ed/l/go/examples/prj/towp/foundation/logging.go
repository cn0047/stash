package foundation

import (
	"github.com/to-com/go-log/zapx"
	"go.uber.org/zap"
)

func NewLogger() *zap.SugaredLogger {
	logger := zapx.NewWithServiceName("wp")
	return logger
}
