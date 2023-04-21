package internal

import (
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/internal/business"
	"github.com/to-com/wp/internal/pubsub"
	"github.com/to-com/wp/internal/repository"
	"github.com/to-com/wp/internal/service/auth"
	"github.com/to-com/wp/internal/service/tsc"
	"go.uber.org/zap"
)

type Application struct {
	cfg     *config.Config
	logger  *zap.SugaredLogger
	handler *HTTPHandler
}

func New(logger *zap.SugaredLogger) (*Application, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	rep, err := repository.New(cfg, logger)
	if err != nil {
		return nil, err
	}
	pubSubClient, err := pubsub.New(cfg, logger)
	if err != nil {
		return nil, err
	}

	httpClient := NewHTTPClient(cfg)
	serviceCatalogService := tsc.New(cfg, logger, httpClient)

	bs := business.New(cfg, logger, rep, pubSubClient, serviceCatalogService)

	authService := auth.New(cfg, logger, httpClient)
	handler := NewHTTPHandler(logger, bs, authService)
	return &Application{
		cfg:     cfg,
		logger:  logger,
		handler: handler,
	}, nil
}
