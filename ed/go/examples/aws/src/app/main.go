//
// docker run -p 8000:8000 amazon/dynamodb-local
//
package main

import (
	"app/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"

	"app/dynamodb"
)

var (
	aws_access_key_id     = ""
	aws_secret_access_key = ""
	region                = "us-east-1"
	bucket                = ""

	host = "http://127.0.0.1:8000"
)

func main() {
	c1 := GetAWSConfig2()
	c2 := GetAWSConfig2()
	s3.Run(c1, bucket)
	return
	dynamodb.Run(c2)
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
