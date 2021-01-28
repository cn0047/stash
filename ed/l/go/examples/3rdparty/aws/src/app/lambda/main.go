package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/thepkg/awsl"

	"app/internal"
	"app/s3"
	"app/sts"
)

const insert = "INSERT"

func main() {
	lambda.Start(Handler4)
}

func Handler4(event events.DynamoDBEvent) {
	cfg := internal.GetAWSConfig()
	bucket := ""
	roleToAssume := ""

	r, err := sts.AssumeRole(cfg, roleToAssume)
	if err != nil {
		log.Fatalf("err1: %#v \n", err)
	}

	c, err := internal.GetConfigByCredentials(
		"us-east-1",
		*r.Credentials.AccessKeyId,
		*r.Credentials.SecretAccessKey,
		*r.Credentials.SessionToken,
	)
	if err != nil {
		log.Fatalf("err2: %#v \n", err)
	}

	s3.PutToS3(c, bucket, "/m.txt", strings.NewReader("it works!!"))
}

func Handler3(event events.DynamoDBEvent) {
	rtl(event.Records)
	for _, record := range event.Records {
		rtl(record)
		r := awsl.FromDynamoDBMap(record.Change.NewImage)
		rtl(r)
	}
}

func Handler2(event events.DynamoDBEvent) {
	data := make([]byte, 0, 0)
	rtl(event.Records)
	for _, record := range event.Records {
		r := awsl.FromDynamoDBMap(record.Change.NewImage)
		r["updated_at"] = r["created_at"]
		r["event_name"] = record.EventName

		payload, err := json.Marshal(r)
		if err != nil {
			panic(err)
		}
		data = append(data, payload...)
	}
	rtl(map[string]string{"data": strings.Replace(fmt.Sprintf("%s", data), `"`, "", -1)})

	cfg := internal.GetAWSConfig()
	res := s3.PutToS3(cfg, internal.GetBucket(), getS3Key(), bytes.NewReader(data))
	rtl(res)
}

func getS3Key() string {
	t := time.Now()
	s := fmt.Sprintf(
		"/x-%d-%d-%d_%d-%d-%d_%d.json",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(),
	)

	return s
}

func Handler1(event events.DynamoDBEvent) {
	log.Printf("\n Got: %#v", event)
	/*
	   Got: events.DynamoDBEvent{
	   Records:[]events.DynamoDBEventRecord{
	       events.DynamoDBEventRecord{
	           EventName:"INSERT|MODIFY|REMOVE",
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

	// Just to test config.
	if !config.UseRealTimeLog {
		rtl("epic fail")
		return
	}

	// Post event to RealTimeLog.
	rtl(event)

	for _, record := range event.Records {
		key := record.Change.NewImage["key"].String()
		val := record.Change.NewImage["val"].String()
		valInt64, err := record.Change.NewImage["val_int"].Integer()
		if err != nil {
			panic(fmt.Errorf("ERR-3: %w", err))
		}
		nonExistingValue := int64(-1)
		nonExistingValueErr := fmt.Errorf("ERR: non_existing_value_err")
		v, ok := record.Change.NewImage["non_existing_value"]
		if ok {
			nonExistingValue, nonExistingValueErr = v.Integer()
		}
		data := map[string]interface{}{
			"eID":                    record.EventID,
			"e":                      record.EventName,
			"k":                      key,
			"v":                      val,
			"vi":                     valInt64,
			"non_existing_value":     nonExistingValue,
			"non_existing_value_err": nonExistingValueErr,
		}
		rtl(data)
	}
}

func rtl(data interface{}) {
	j, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Errorf("ERR-RTL-1: %w", err))
	}
	_, err2 := http.Post("https://realtimelog.herokuapp.com:443/64kfym341kp2", "application/json", bytes.NewBuffer(j))
	if err2 != nil {
		panic(fmt.Errorf("ERR-RTL-2: %w", err2))
	}
}
