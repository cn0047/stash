package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	"three/client"
	"three/proto/backend/proto"
	"three/server"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("provide param server or client")
	}

	addr := ":50053"

	switch os.Args[1] {
	case "server":
		startServer(addr)
	case "client":
		testClient(addr)
	}
}

func testClient(addr string) {
	conn, c := client.NewMainServerClient(addr)
	ping(c)
	echo(c)

	err := conn.Close()
	if err != nil {
		log.Fatalf("failed to close client conn, err: %#v", err)
	}
}

func echo(client proto.MainServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Echo(ctx, &proto.EchoRequest{Message: "ping"})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Unimplemented:
				fmt.Printf("GRPC err: %+v \n", e.Message())
			case codes.Internal:
				fmt.Printf("GRPC err: InternalError \n")
			default:
				fmt.Printf("GRPC err: %+v, %+v \n", e.Code(), e.Message())
			}
		}
		log.Fatalf("failed to perform echo, res: %#v, err: %#v", r, err)
	}

	log.Printf("\nRes: %#v \nErr: %#v", r, err)
}

func ping(client proto.MainServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.Ping(ctx, &proto.PingRequest{Message: "ping"})
	if err != nil {
		log.Fatalf("failed to perform ping, res: %#v, err: %#v", r, err)
	}

	log.Printf("\nRes: %#v \nErr: %#v", r, err)
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen net, err: %#v", err)
	}

	s := grpc.NewServer()
	mainServer := &server.MainServer{}
	proto.RegisterMainServiceServer(s, mainServer)
	reflection.Register(s)
	fmt.Printf("Starting to serve on: %s\n", addr)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve, err: %#v", err)
	}
}
