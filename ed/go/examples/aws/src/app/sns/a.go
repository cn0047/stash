package sns

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func Run(cfg *aws.Config) {
	svc := sns.New(session.New(), cfg)
	arn := "arn:aws:sns:us-east-1:613225557329:stark-qa-neighbor-onboarding"
	publish(svc, arn)
}

func publish(svc *sns.SNS, arn string) {
	payload := map[string]string{"foo": "bar2"}
	j, err := json.Marshal(payload)
	if err != nil {
		panic(fmt.Sprintf("json err: %v", err))
	}

	input := &sns.PublishInput{
		Message: aws.String(string(j)),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"event_name": {DataType: aws.String("String"), StringValue: aws.String("EVENT_TYPE_FOO")},
		},
		TopicArn: aws.String(arn),
	}
	result, err := svc.Publish(input)
	if err != nil {
		fmt.Println("Publish error:", err)
		return
	}

	fmt.Println(result)
}
