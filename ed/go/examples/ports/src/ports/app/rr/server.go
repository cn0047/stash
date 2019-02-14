package rr

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"ports/app/proto"
	"ports/ddd/application/service"
	"ports/ddd/domain/model"
)

type Server struct{}

func (s *Server) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	res := proto.GetResponse{}
	p, err := service.GetPort(req.ID)
	if err != nil {
		return &res, fmt.Errorf("got get port error: %#v", err)
	}

	res.Data = &proto.Port{
		ID:          p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       p.Alias,
		Regions:     p.Regions,
		Coordinates: p.Coordinates,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}

	return &res, nil
}

func (s *Server) Put(ctx context.Context, req *proto.Port) (*proto.PutResponse, error) {
	var err error
	res := proto.PutResponse{}

	port := model.PortEntity{
		ID:          req.ID,
		Name:        req.Name,
		City:        req.City,
		Country:     req.Country,
		Alias:       req.Alias,
		Regions:     req.Regions,
		Coordinates: req.Coordinates,
		Province:    req.Province,
		Timezone:    req.Timezone,
		Unlocs:      req.Unlocs,
		Code:        req.Code,
	}

	err = service.PutPort(port)
	if err != nil {
		e := fmt.Errorf("failed to perform put port, error: %#v", err)
		res.Error = e.Error()
		return &res, e
	}

	res.Data = "OK"
	log.Printf("stored into db port: %v", port)

	return &res, nil
}

func StartServer() {
	var err error
	lis, err := net.Listen("tcp", RPCPort)
	if err != nil {
		log.Fatalf("failed to listen tcp port: %s, error: %#v", RPCPort, err)
	}

	s := grpc.NewServer()
	proto.RegisterPortServiceServer(s, &Server{})
	reflection.Register(s)

	log.Printf("start to listen tcp port: %s", RPCPort)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to perform serve: %#v", err)
	}
}
