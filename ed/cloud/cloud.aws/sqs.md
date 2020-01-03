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
