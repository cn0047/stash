package one

import (
	"context"
	"fmt"
	"time"
)

type instrumentingMiddleware struct {
	next Service
}

func NewInstrumentingMiddleware(svc Service) Service {
	return instrumentingMiddleware{
		next: svc,
	}
}

func (mw instrumentingMiddleware) Hello(ctx context.Context, key string) (interface{}, error) {
	startedAt := time.Now()
	res, err := mw.next.Hello(ctx, key)
	finishedAt := time.Now()

	fmt.Printf("took: %#v \n", finishedAt.Sub(startedAt))

	return res, err
}
