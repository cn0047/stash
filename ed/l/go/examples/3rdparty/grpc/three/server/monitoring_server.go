package server

import (
	"context"
	"log"

	"three/proto/backend/proto"
)

type MonitoringServer struct {
}

func (m *MonitoringServer) Log(ctx context.Context, in *proto.LogRequest) (*proto.LogResponse, error) {
	log.Print("REG [monitoring]: log")
	res := &proto.LogResponse{Message: "logged", Error: nil}

	return res, nil
}

func (m *MonitoringServer) Health(ctx context.Context, in *proto.HealthRequest) (*proto.HealthResponse, error) {
	return nil, nil
}
