package s3

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Run(cfg *aws.Config, bucket string) {
	f1(cfg, bucket)
}

func f1(cfg *aws.Config, bucket string) {
	svc := s3.New(session.New(), cfg)

	file, err := os.Open(os.Getenv("GOPATH")+"/src/app/s3/s.png")
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
	path := "/test-files/s.png"
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
