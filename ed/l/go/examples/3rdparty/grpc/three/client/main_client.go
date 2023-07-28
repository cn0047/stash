package client

import (
	"log"

	"google.golang.org/grpc"

	"three/proto/backend/proto"
)

func NewMainServerClient(addr string) (*grpc.ClientConn, proto.MainServiceClient) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial net, err: %+v", err)
	}

	client := proto.NewMainServiceClient(conn)

	return conn, client
}
