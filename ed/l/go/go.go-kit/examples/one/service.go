package one

import (
	"context"
	"errors"
)

type Service interface {
	Hello(ctx context.Context, param string) (interface{}, error)
}

type service struct {
}

func NewService() Service {
	return &service{}
}

func (s *service) Hello(_ context.Context, param string) (interface{}, error) {
	if param == "err" {
		return nil, errors.New("some error")
	}
	return "ok", nil
}
