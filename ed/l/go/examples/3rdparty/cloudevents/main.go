package main

import (
	"context"
	"fmt"
	"log"
	"time"

	ce "github.com/cloudevents/sdk-go/v2"
	cec "github.com/cloudevents/sdk-go/v2/client"
)

const (
	URI = "http://localhost:8080"
)

func main() {
	var err error
	//err = example1()
	err = example2()
	if err != nil {
		log.Fatalf("err: %v", err)
	}
}

func getClient() (cec.Client, error) {
	c, err := ce.NewClientHTTP()

	return c, err
}

type Event1 struct {
	Action    string
	CreatedAt time.Time
	Msg       string
}

func example1() error {
	c, err := getClient()
	if err != nil {
		return fmt.Errorf("failed to create client, err: %w", err)
	}

	event := ce.NewEvent()
	event.SetID("ID1")
	event.SetSpecVersion("1.0")
	event.SetType("example.type1")
	event.SetSource("example/source1")
	event.SetTime(time.Now())
	err = event.SetData(ce.ApplicationJSON, map[string]string{"hello": "world"})
	if err != nil {
		return fmt.Errorf("failed to set data, err: %w", err)
	}

	ctx := ce.ContextWithTarget(context.Background(), URI)
	res := c.Send(ctx, event)
	undelivered := ce.IsUndelivered(res)
	fmt.Printf("undelivered: %v, res: %v\n", undelivered, res)

	return nil
}

func example2() error {
	c, err := getClient()
	if err != nil {
		return fmt.Errorf("failed to create client, err: %w", err)
	}

	event := ce.NewEvent()
	event.SetID("ID1")
	event.SetSpecVersion("1.0")
	event.SetType("example.type1")
	event.SetSource("example/source1")
	event.SetTime(time.Now())
	payload := Event1{
		Action:    "test",
		CreatedAt: time.Now(),
		Msg:       "example2",
	}
	err = event.SetData(ce.ApplicationJSON, payload)
	if err != nil {
		return fmt.Errorf("failed to set data, err: %w", err)
	}

	ctx := ce.ContextWithTarget(context.Background(), URI)
	res := c.Send(ctx, event)
	undelivered := ce.IsUndelivered(res)
	fmt.Printf("undelivered: %v, res: %v\n", undelivered, res)

	return nil
}
