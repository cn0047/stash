package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"app/lib"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := lib.NewOneClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Simple(ctx, &lib.Request{Msg: "hey"})
	if err != nil {
		log.Fatalf("fail, error: %v", err)
	}

	log.Printf("Ok: %s", r.Msg)

	er2 := conn.Close()
	if er2 != nil {
		log.Fatalf("failed to close conn, error: %v", er2)
	}
}
