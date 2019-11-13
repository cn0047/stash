//
// docker run -p 8000:8000 amazon/dynamodb-local
//
package main

import (
	"app/dynamodb"
	"app/s3"
	"app/sqs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var (
	aws_access_key_id     = ""
	aws_secret_access_key = ""
	region                = "eu-central-1"
	bucket                = "basicbkt"

	host = "http://127.0.0.1:8000"
)

func main() {
	c := GetAWSConfig()
	sqs.Run(c)
	return
	s3.Run(c, bucket)
	dynamodb.Run(c)
}

func GetAWSConfig2() *aws.Config {
	creds := credentials.NewStaticCredentials("123", "123", "")
	awsConfig := &aws.Config{
		Credentials: creds,
		Region:      &region,
		Endpoint:    &host,
	}

	return awsConfig
}

func GetAWSConfig() *aws.Config {
	token := ""

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	return cfg
}
