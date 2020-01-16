package s3

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"app/sts"
)

func Run(cfg *aws.Config, bucket string) {
	//f1(cfg, bucket)
	//f3(cfg)

	PutToS3(cfg, bucket, "/msg.txt", strings.NewReader("it works!"))

	//_, _, fileBytes := getFileBytes()
	//PutToS3(cfg, bucket, "/test.png", fileBytes)
}

func f3(cfg *aws.Config) {
	r, err := sts.AssumeRole(cfg, "arn:aws:iam::${id}:role/${role}")
	if err != nil {
		log.Fatalf("failed to perform sts.AssumeRole, error: %v \n", err)
	}
	fmt.Printf(
		"STS Credentials: id: %v, key: %v, tkn: %v",
		r.Credentials.AccessKeyId,
		r.Credentials.SecretAccessKey,
		r.Credentials.SessionToken,
	)
}

func PutToS3(cfg *aws.Config, bucket string, key string, body io.ReadSeeker) interface{} {
	svc := s3.New(session.New(), cfg)

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	}
	res, err := svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func f2(cfg *aws.Config, bucket string) {
	key := "/test-files/s.png"
	input := s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
	}

	obj := s3manager.BatchUploadObject{
		Object: &input,
	}

	_ = obj
}

func getFileBytes() (int64, string, *bytes.Reader) {
	file, err := os.Open(os.Getenv("GOPATH") + "/src/app/s3/s.png")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	size := fileInfo.Size()
	buffer := make([]byte, size)

	_, err = file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	return size, fileType, fileBytes
}

func f1(cfg *aws.Config, bucket string) {
	svc := s3.New(session.New(), cfg)

	size, fileType, fileBytes := getFileBytes()
	path := "/test-files/s.png"
	arns := "arn:aws:s3:::basicbkt"

	apARN, err := arn.Parse(arns)
	if err != nil {
		log.Fatal(err)
	}

	params := &s3.PutObjectInput{
		//Bucket:        aws.String(bucket),
		Bucket:        aws.String(apARN.Resource),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	resp, err := svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response %s", awsutil.StringValue(resp))
}
