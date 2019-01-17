package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"app/lib"
)

type server struct{}

func (s *server) Simple(ctx context.Context, in *lib.Request) (*lib.Response, error) {
	log.Printf("Got: %v", in.Msg)
	return &lib.Response{Msg: "Copy " + in.Msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	lib.RegisterOneServer(s, &server{})
	reflection.Register(s)

	er2 := s.Serve(lis)
	if er2 != nil {
		log.Fatalf("failed to serve: %v", er2)
	}
}
