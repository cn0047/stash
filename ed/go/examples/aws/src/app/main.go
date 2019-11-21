//
// docker run -p 8000:8000 amazon/dynamodb-local
//
package main

import (
	"app/dynamodb"
	"app/internal"
	"app/s3"
	"app/sqs"
)

func main() {
	c := internal.GetAWSConfig()
	dynamodb.Run(c)
	return
	sqs.Run(c)
	s3.Run(c, internal.GetBucket())
}
