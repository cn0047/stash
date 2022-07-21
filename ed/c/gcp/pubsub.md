Pub/Sub
-

````sh
gcloud pubsub topics list
gcloud pubsub topics publish $topic --message='{"type":"test","data":"ok"}'

gcloud pubsub subscriptions list

gcloud pubsub subscriptions pull --auto-ack --limit=1 $subscription
````

Subscription delivery types: pull, push.
