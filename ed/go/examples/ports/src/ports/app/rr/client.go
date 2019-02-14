package rr

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"

	"ports/app/proto"
	"ports/ddd/application/service"
	"ports/ddd/domain/model"
)

var (
	conn *grpc.ClientConn
)

func StartClient() {
	StartRPC()
	InitDB()
	StartREST()
}

func StartRPC() {
	var err error

	conn, err = grpc.Dial(RPCHost+RPCPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect, error: %#v", err)
	}

	log.Printf("rpc connection (%s%s) established", RPCHost, RPCPort)
}

func CloseConn() {
	err := conn.Close()
	if err != nil {
		log.Fatalf("failed to close conn, error: %v", err)
	}
}

func InitDB() {
	err := service.LoadPortsFromJSONFile(PortsDataFile, func(port model.PortEntity) error {
		// todo: implement batching.
		return PutPort(port)
	})
	if err != nil {
		log.Fatalf("failed to open file for init DB, error: %#v", err)
	}
}

func PutPort(port model.PortEntity) error {
	var err error

	c := proto.NewPortServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.Port{
		ID:          port.ID,
		Name:        port.Name,
		City:        port.City,
		Country:     port.Country,
		Alias:       port.Alias,
		Regions:     port.Regions,
		Coordinates: port.Coordinates,
		Province:    port.Province,
		Timezone:    port.Timezone,
		Unlocs:      port.Unlocs,
		Code:        port.Code,
	}

	_, err = c.Put(ctx, &req)
	if err != nil {
		log.Fatalf("fail to perform RPC put, error: %#v", err)
	}

	return nil
}

func GetPort(id string) (port model.PortEntity, err error) {
	c := proto.NewPortServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := proto.GetRequest{ID: id}
	res, err := c.Get(ctx, &req)
	if err != nil {
		return port, fmt.Errorf("fail to perform RPC get, error: %#v", err)
	}
	if res.Error != "" {
		return port, fmt.Errorf("got RPC error: %s", res.Error)
	}

	port = model.PortEntity{
		ID:          res.Data.ID,
		Name:        res.Data.Name,
		City:        res.Data.City,
		Country:     res.Data.Country,
		Alias:       res.Data.Alias,
		Regions:     res.Data.Regions,
		Coordinates: res.Data.Coordinates,
		Province:    res.Data.Province,
		Timezone:    res.Data.Timezone,
		Unlocs:      res.Data.Unlocs,
		Code:        res.Data.Code,
	}

	return port, nil
}
