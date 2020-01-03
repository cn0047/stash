//
// docker run -p 8000:8000 amazon/dynamodb-local
//
package main

import (
	"app/dynamodb"
	"app/internal"
	"app/s3"
	"app/sns"
	"app/sqs"
	"app/ssm"
)

func main() {
	c := internal.GetAWSConfig()
	s3.Run(c, internal.GetBucket())
	return
	sns.Run(c)
	ssm.Run(c)
	dynamodb.Run(c)
	sqs.Run(c)
}
