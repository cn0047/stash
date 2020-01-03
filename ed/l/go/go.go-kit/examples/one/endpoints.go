package one

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeHelloEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return svc.Hello(ctx, request.(string))
	}
}
