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
