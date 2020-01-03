package sqs

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	QueueName = "simpleQueue"
)

var (
	QueueUrl = "https://sqs.eu-central-1.amazonaws.com/915063445062/" + QueueName
)

func Run(cfg *aws.Config) {
	svc := sqs.New(session.New(), cfg)
	add(svc)
	get(svc)
}

func getQueueUrl(svc *sqs.SQS) (queueName *string) {
	out, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(QueueName),
	})
	if err != nil {
		log.Printf("failed to set queue, error: %v \n", err)
	}

	return out.QueueUrl
}

func add(svc *sqs.SQS) {
	msg := sqs.SendMessageInput{
		QueueUrl:               aws.String(QueueUrl),
		MessageBody:            aws.String(`{"n":200}`),
		//MessageDeduplicationId: aws.String("200"), // only for FIFO queue.
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"code": {DataType: aws.String("Number"), StringValue: aws.String("100")},
		},
	}
	out, err := svc.SendMessage(&msg)
	if err != nil {
		log.Printf("failed to add to queue, error: %v \n", err)
	}

	log.Printf("Message queued: %v \n", out)
}

func render(res *sqs.ReceiveMessageOutput) {
	for _, m := range res.Messages {
		d := make(map[string]int, 1)
		err := json.Unmarshal([]byte(*m.Body), &d)
		if err != nil {
			log.Printf("failed to unmarshal body, error: %v \n", err)
		}

		log.Printf(
			"got message, id: %v, code: %v, data: %v \n",
			*m.MessageId,
			*m.MessageAttributes["code"].StringValue,
			d,
		)
	}
}

func get(svc *sqs.SQS) {
	res, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: getQueueUrl(svc),
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: aws.StringSlice([]string{"All"}),
		WaitTimeSeconds: aws.Int64(1), // 1 second
	})
	if err != nil {
		log.Printf("failed to get from queue, error: %v \n", err)
	}

	render(res)
	if len(res.Messages) > 0 {
		del(svc, res.Messages[0].ReceiptHandle)
	}
}

func del(svc *sqs.SQS, identifier *string) {
	res, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl: getQueueUrl(svc),
		ReceiptHandle: identifier,
	})
	if err != nil {
		log.Printf("failed to delete from queue, error: %v \n", err)
	}

	log.Printf("Message deleted: %v \n", res)
}
