// @example:
// go run main.go info s1 ''
// go run main.go consumeFromConsumer 's1' 'jetStream.test.simple' 'c1'
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	ServerURL           = "nats://0.0.0.0:53341"
	UserName            = "local"
	UserPassword        = ""
	DefaultStreamName   = "s1"
	DefaultSubjectName  = "jetStream.test.simple"
	DefaultConsumerName = "c1"

	key = "[v1] "
)

type Message struct {
	Text string
}

func main() {
	action := os.Args[1]
	streamName := os.Args[2]
	subjectName := os.Args[3]
	consumerName := os.Args[4]

	if streamName == "" {
		streamName = DefaultStreamName
	}
	if subjectName == "" {
		subjectName = DefaultSubjectName
	}
	if consumerName == "" {
		consumerName = DefaultConsumerName
	}

	c := getConn()
	defer c.Drain()
	js := getJetStreamContext(c)

	switch action {
	case "publish":
		publish(js, subjectName, key+"hello")
	case "publishMany":
		publishMany(js, subjectName, key+"hello")
	case "consume":
		consume(js, subjectName)
	case "consumeFromConsumer":
		consumeFromConsumer(js, streamName, subjectName, consumerName)
	case "info":
		printStreamInfo(js, streamName)
	default:
		printStreamInfo(js, streamName)
	}

	fmt.Printf("done\n")
}

func getConn() *nats.Conn {
	c, err := nats.Connect(ServerURL, nats.UserInfo(UserName, UserPassword))
	e(err)
	return c
}

func getJetStreamContext(c *nats.Conn) nats.JetStreamContext {
	js, err := c.JetStream()
	e(err)
	return js
}

func printStreamInfo(js nats.JetStreamContext, streamName string) {
	stream, err := js.StreamInfo(streamName)
	e(err)
	fmt.Printf("stream: %v\n", stream)
}

func createStream(js nats.JetStreamContext, streamName string, subjectName string) {
	_, err := js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: []string{subjectName},
	})
	e(err)
}

func publish(js nats.JetStreamContext, subjectName string, msg string) {
	data, err := json.Marshal(Message{Text: msg})
	e(err)
	_, err = js.Publish(subjectName, data)
	e(err)
}

func publishMany(js nats.JetStreamContext, subjectName string, msg string) {
	for i := 0; i < 100; i++ {
		time.Sleep(10 * time.Millisecond)
		publish(js, subjectName, fmt.Sprintf("app message %s #%v", msg, i))
		fmt.Printf("published: %v\n", i)
	}
}

func consume(js nats.JetStreamContext, subjectName string) {
	_, err := js.Subscribe(subjectName, func(m *nats.Msg) {
		err := m.Ack()
		e(err)

		data := ""
		var msg Message
		err = json.Unmarshal(m.Data, &msg)
		if err != nil {
			// Fallback to raw string.
			data = string(m.Data)
		}
		fmt.Printf("got: %v\n", data)
	})
	e(err)
}

func consumeFromConsumer(js nats.JetStreamContext, streamName, subjectName, consumerName string) {
	sub, err := js.PullSubscribe(subjectName, consumerName, nats.Bind(streamName, consumerName))
	e(err)

	msgs, err := sub.Fetch(1, nats.MaxWait(5*time.Second))
	e(err)

	for _, msg := range msgs {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		err := msg.Ack()
		e(err)
	}
}

func e(err error) {
	if err != nil {
		panic(err)
	}
}
