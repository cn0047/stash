package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var (
	aws_access_key_id     = ""
	aws_secret_access_key = ""
	region                = "us-east-1" // us-east-1|eu-central-1
	bucket                = "basicbkt"

	host = "http://127.0.0.1:8000"
)

func GetBucket() string {
	return bucket
}

func GetAWSConfig() *aws.Config {
	return getStaticConfig()
}

func getStaticConfig() *aws.Config {
	cfg := &aws.Config{
		Region: &region,
	}

	return cfg
}

func getStaticConfig1() *aws.Config {
	creds := credentials.NewStaticCredentials("123", "123", "")
	cfg := &aws.Config{
		Credentials: creds,
		Region:      &region,
		Endpoint:    &host,
	}

	return cfg
}

func getStaticConfig2() *aws.Config {
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, "")
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)

	return cfg
}
