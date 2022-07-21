SQS (Simple Queue Service)
-

[docs](https://docs.aws.amazon.com/sqs/index.html)

DLQ - Dead-Letter Queue.

Queue types: standard, fifo.

Use to pull message.

Queue Attributes:
* Message Retention Period - days before AWS will delete item from queue.
* Maximum Message Size.
* Delivery Delay - time to delay first delivery.
* Receive Message Wait Time - max amount of time
  that a long polling receive call will wait for a message to become available.
* Visibility Timeout - period of time during which SQS prevents other consumers
  from receiving the message which sent to other consumer but not acknowledged yet.

````sh
aws --profile=$p sqs list-queues

aws sqs get-queue-attributes \
  --queue-url https://sqs.us-west-2.amazonaws.com/$SOME_ID/$SOME_Q \
  --attribute-names ApproximateNumberOfMessages ApproximateNumberOfMessagesDelayed ApproximateNumberOfMessagesNotVisible
  
````
