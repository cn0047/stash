package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"three/proto/backend/proto"
)

var (
	ErrCodeNotImplemented int32 = 1
)

type MainServer struct {
}

func (s *MainServer) Health(ctx context.Context, req *proto.HealthRequest) (*proto.HealthResponse, error) {
	log.Print("REG [main]: health")

	res := &proto.HealthResponse{Message: "healthy", Error: &proto.AppError{ErrorCode: 0, ErrorMessage: "_"}}

	return res, nil
}

func (s *MainServer) Ping(ctx context.Context, req *proto.PingRequest) (*proto.PingResponse, error) {
	log.Print("REG [main]: ping")

	res := &proto.PingResponse{Message: "pong", Error: nil}

	return res, nil
}

func (*MainServer) Echo(ctx context.Context, req *proto.EchoRequest) (*proto.EchoResponse, error) {
	log.Print("REG [main]: echo")

	res := &proto.EchoResponse{
		Message: "method Echo not implemented yet",
		Error: &proto.AppError{
			ErrorCode:    ErrCodeNotImplemented,
			ErrorMessage: fmt.Sprintf("AppErrCode: %d", ErrCodeNotImplemented),
		},
	}

	return res, status.Errorf(codes.Unimplemented, res.Message)
}

func (s *MainServer) Gossip(req *proto.GossipRequest, stream proto.MainService_GossipServer) error {
	log.Print("REG [main]: gossip")

	for i := 1; i <= 20; i++ {
		time.Sleep(900 * time.Millisecond)
		msg := fmt.Sprintf("%s - %v", req.Message, i)
		res := &proto.GossipResponse{Message: msg}
		if err := stream.Send(res); err != nil {
			return fmt.Errorf("failed to send stream data, err: %w", err)
		}
		log.Printf("sent: %s", msg)
	}

	return nil
}
