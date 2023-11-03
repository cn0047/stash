package main

import (
	"context"
	"fmt"
	"io"
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
	if len(os.Args) <= 1 {
		log.Fatalf("provide param server or client")
	}

	addr := ":50053"

	switch os.Args[1] {
	case "server":
		startServer(addr)
	case "client":
		testClient(addr, os.Args[1:])
	}
}

func testClient(addr string, params []string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	connMain, cMain := client.NewMainServerClient(addr)
	connMon, cMon := client.NewMonitoringServerClient(addr)

	if len(params) > 1 {
		switch params[1] {
		case "health":
			health(ctx, cMain)
		case "ping":
			ping(ctx, cMain)
		case "echo":
			echo(ctx, cMain)
		case "gossip":
			gossip(ctx, cMain)
		case "log":
			mLog(ctx, cMon)
		default:
			ping(ctx, cMain)
		}
	} else {
		ping(ctx, cMain)
	}

	if err := connMain.Close(); err != nil {
		log.Fatalf("failed to close connMain, err: %#v", err)
	}
	if err := connMon.Close(); err != nil {
		log.Fatalf("failed to close connMon, err: %#v", err)
	}
}

func health(ctx context.Context, client proto.MainServiceClient) {
	r, err := client.Health(ctx, &proto.HealthRequest{})
	if err != nil {
		log.Fatalf("failed to perform health, res: %#v, err: %#v", r, err)
	}

	log.Printf("\nRes: %#v \nErr: %#v", r, err)
}

func echo(ctx context.Context, client proto.MainServiceClient) {
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

func ping(ctx context.Context, client proto.MainServiceClient) {
	r, err := client.Ping(ctx, &proto.PingRequest{Message: "ping"})
	if err != nil {
		log.Fatalf("failed to perform ping, res: %#v, err: %#v", r, err)
	}

	log.Printf("\nRes: %#v \nErr: %#v", r, err)
}

func gossip(ctx context.Context, client proto.MainServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	r, err := client.Gossip(ctx, &proto.GossipRequest{Message: "rumor"})
	if err != nil {
		log.Fatalf("failed to perform gossip, res: %#v, err: %#v", r, err)
	}

	for {
		data, err := r.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to get stream data, err: %#v", err.Error())
		}
		log.Printf("\nRes chunk: %#v", data.Message)
	}

	log.Printf("\nRes: end")
}

func mLog(ctx context.Context, client proto.MonitoringServiceClient) {
	r, err := client.Log(ctx, &proto.LogRequest{Message: "log"})
	if err != nil {
		log.Fatalf("failed to perform log, res: %#v, err: %#v", r, err)
	}

	log.Printf("\nRes: %#v \nErr: %#v", r, err)
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen net, err: %#v", err)
	}

	s := grpc.NewServer()
	proto.RegisterMainServiceServer(s, &server.MainServer{})
	proto.RegisterMonitoringServiceServer(s, &server.MonitoringServer{})
	reflection.Register(s)
	fmt.Printf("Starting to serve on: %s\n", addr)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve, err: %#v", err)
	}
}
