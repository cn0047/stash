SNS (Simple Notification Service)
-

[docs](https://docs.aws.amazon.com/sns/)

SNS - pub/sub (many-to-many, fan out) messaging service.

Use to push message.

Types of clients:
* Publisher (aka Producer)
* Subscriber (aka Consumer)

The type of endpoint to subscribe:
* HTTP
* HTTPS
* Email
* Email-Json
* AWS SQS
* AWS Lambda
* Platform Application Endpoint
* SMS

Components:
* Topic
* Subscription
* Publisher

Topic settings:
* Delivery retry policy:
  * 3 - default number of retries.
  * 20 seconds - default minimum delay.
  * 20 seconds - default max delay.
  * Linear - default retry-backoff function.
* Delivery status logging.

Subscription settings:
* Subscription filter policy.
* Redrive policy (dead-letter queue):
  * dead-letter queue disabled by default.
* Delivery retry policy.
  * 3 - default number of retries.
  * 20 seconds - default minimum delay.
  * Linear - default retry-backoff function.

````sh
# create topic
aws sns create-topic --name=mytopic
````

Subscription filter policy:
````
{
  "event_name": ["MY_EVENT_NAME"]
}
````
