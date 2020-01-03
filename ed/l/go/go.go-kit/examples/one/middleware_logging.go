package one

import (
	"context"
	"log"
)

type loggingMiddleware struct {
	next Service
}

func NewLoggingMiddleware(svc Service) Service {
	return loggingMiddleware{
		next: svc,
	}
}

func (mw loggingMiddleware) Hello(ctx context.Context, key string) (interface{}, error) {
	res, err := mw.next.Hello(ctx, key)
	if err != nil {
		log.Printf("error: %#v", err)
	}

	return res, err
}
