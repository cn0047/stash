//
// docker run -p 8000:8000 amazon/dynamodb-local
//
package main

import (
	"app/dynamodb"
	"app/internal"
	"app/s3"
	"app/sqs"
	"app/ssm"
)

func main() {
	c := internal.GetAWSConfig()
	ssm.Run(c)
	return
	dynamodb.Run(c)
	sqs.Run(c)
	s3.Run(c, internal.GetBucket())
}
