package server

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"three/proto/backend/proto"
)

var (
	ErrCodeNotImplemented int32 = 1
)

type MainServer struct {
}

func (s *MainServer) Ping(ctx context.Context, request *proto.PingRequest) (*proto.PingResponse, error) {
	log.Print("REG: ping")
	res := &proto.PingResponse{Message: "pong", Error: nil}

	return res, nil
}

func (*MainServer) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	log.Print("REG: echo")
	res := &proto.EchoResponse{
		Message: "method Echo not implemented yet",
		Error: &proto.AppError{
			ErrorCode:    ErrCodeNotImplemented,
			ErrorMessage: fmt.Sprintf("AppErrCode: %d", ErrCodeNotImplemented),
		},
	}

	return res, status.Errorf(codes.Unimplemented, res.Message)
}
