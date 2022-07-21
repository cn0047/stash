Kinesis
-

* Stream.
* Firehose.
* Analytics.

Stream:
````
EC2/Servers/Mobile clients -> Kinesis streams -> App/Lambda -> s3/DynamoDB/RedShift
````

Read limitations:
1 shard can return <2MB of data per second.
5 reads per shard per second.
`GetRecords` can return <10MB.

````sh
aws kinesis list-streams
aws kinesis describe-stream
aws kinesis describe-stream-summary

aws kinesis list-shards

aws firehose list-delivery-streams
aws firehose put-record
````
