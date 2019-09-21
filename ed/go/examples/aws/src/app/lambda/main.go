package main

import (
	"bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/http"
)

const insert = "INSERT"

//Handler process lambda event
func Handler(event events.DynamoDBEvent) {
	log.Printf("\n Got: %#v", event)
	/*
	   Got: events.DynamoDBEvent{
	   Records:[]events.DynamoDBEventRecord{
	       events.DynamoDBEventRecord{
	           EventName:"INSERT",
	           EventID:"fdf85492723daa58508d6b9a8731f45d",
	           Change:events.DynamoDBStreamRecord{
	               ApproximateCreationDateTime:events.SecondsEpochTime{Time:time.Time{wall:0x0, ext:63704222878, loc:(*time.Location)(0xb05020)}},
	               Keys:map[string]events.DynamoDBAttributeValue{"id":events.DynamoDBAttributeValue{value:"i32", dataType:8}, "key":events.DynamoDBAttributeValue{value:"i32", dataType:8}},
	               NewImage:map[string]events.DynamoDBAttributeValue{"id":events.DynamoDBAttributeValue{value:"i32", dataType:8}, "key":events.DynamoDBAttributeValue{value:"i32", dataType:8}},
	               OldImage:map[string]events.DynamoDBAttributeValue(nil),
	               SequenceNumber:"587300000000003193252793", SizeBytes:22, StreamViewType:"NEW_AND_OLD_IMAGES"
	           },
	           AWSRegion:"us-west-2",
	           EventSource:"aws:dynamodb", EventVersion:"1.1", EventSourceArn:"arn:aws:dynamodb:us-west-2:613225557329:table/KovpakTest/stream/2019-09-16T09:19:17.215", UserIdentity:(*events.DynamoDBUserIdentity)(nil)}
	       }
	   }
	*/

	// Post event to realtimelog.
	j, _ := json.Marshal(event)
	http.Post("https://realtimelog.herokuapp.com:443/64kfym341kp", "application/json", bytes.NewBuffer(j))
}

func main() {
	lambda.Start(Handler)
}
