package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

var (
	aws_access_key_id     = ""
	aws_secret_access_key = ""
	region                = "us-east-1"
	bucket                = ""
)

func main() {
	v1()
}

func v1() {
	token := ""

	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		panic(err)
	}
	cfg := aws.NewConfig().WithRegion(region).WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open("s.png")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/legacy-files/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		panic(err)
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
}
