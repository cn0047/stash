DynamoDB
-

[docs](https://docs.aws.amazon.com/dynamodb/?id=docs_gateway)
[examples](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/using-dynamodb-with-go-sdk.html)
[limits](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html)

````
Relational DB ⇒ Table ⇒ Row  ⇒ Column    ⇒ Primary key ⇒ View
DDB           ⇒ Table ⇒ Item ⇒ Attribute ⇒ Primary key ⇒ Global Secondary Index
````

EventName: `INSERT|MODIFY|REMOVE`
StreamViewType: `KEYS_ONLY|NEW_IMAGE|OLD_IMAGE|NEW_AND_OLD_IMAGES`

Table - reqular table.
<br>Global Table - multi-region, and multi-master.

Table may have TTL attribute.

A Scan operation (In general less efficient than other operations)
reads every item in a table or a secondary index.
By default Scan returns all of the data attributes for every item in the table or index.
Scan always returns a result set.

You can define a maximum of 5 local secondary indexes
and 20 global secondary indexes per table.

PROVISIONED - Sets the read/write capacity mode to PROVISIONED.
PAY_PER_REQUEST - for unpredictable workloads.

DynamoDB trigger - is lambda function.

Read Capacity Units (RCUs) = Round up(Item size/4KB)/2.

````js
// item to add in AWS DDB console:
{
  "bin": {"B": "AAEqQQ=="},
  "binset": {"BS": ["AAEqQQ==", "AAEqQQ=="] },
  "bool": {"BOOL": false },
  "float": {"N": "0.9998"},
  "int": {"N": "-256"},
  "key": "k5",
  "list": {"L": [{"S": "1"}, {"BOOL": true }, {"NULL": true } ] },
  "map": {"M": {"map": {"S": "str"}, "map2": {"M": {"100": {"N": "STR MUST NOT BE BLANK"} } } } },
  "null": {"NULL": true },
  "numset": {"NS": ["-99", "99"] },
  "numset2map": {"M": {"test": {"BOOL": true } } },
  "str": {"S": "str"},
  "strset": {"SS": ["foo", "bar"] },
  "strset2map": {"M": {"foo": {"S": "bar"} } },
  "strset3list": {"L": [{"S": "foo"}, {"S": "bar"} ] }
}
````

````sh
tbl=hotdata

aws --profile=$p dynamodb list-global-tables
aws --profile=$p dynamodb list-tables

# ONLY PRIMARY KEY required all other fields may be omitted
aws dynamodb create-table \
  --table-name $tbl \
  --billing-mode=PAY_PER_REQUEST \
  --attribute-definitions AttributeName=user_id,AttributeType=N \
  --key-schema AttributeName=user_id,KeyType=HASH \

# count
aws dynamodb scan --table-name $tbl | jq -c '.Count'

aws dynamodb scan --table-name $tbl \
  --filter-expression "owner_id=:oId AND d_id=:dId" \
  --expression-attribute-values '{":oId":{"N":"1"},":dId":{"N":"2"}}'
# result: {"Items": [...], "Count": 45, "ScannedCount": 1291}

# search by primary key
aws dynamodb query --table-name $tbl \
  --key-condition-expression "key_id=:kId" \
  --expression-attribute-values  '{":kId":{"S":"3237415"}}' \
  | jq -c '.Count,.ScannedCount'

# add new item
aws --profile=$p  dynamodb put-item --table-name $tbl \
  --item '{"key":{"S": "k1"},"v":{"S": "v1"}}'

aws dynamodb put-item --table-name $tbl --item '{
  "email"      : {"S": "x@y.com"},
  "err"        : {"BOOL": false},
  "created_at" : {"S": "2019-10-31T16:32:39.243443+02:00"},
  "type"       : {"S": "test"}
}'

aws dynamodb get-item --table-name=$tbl --key='{"user_id": {"S": "'$uId'"}}'
````
