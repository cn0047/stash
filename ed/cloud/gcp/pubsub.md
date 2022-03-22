Pub/Sub
-

````sh
gcloud pubsub topics list
gcloud pubsub topics publish $topic --message='{"type":"test","data":"ok"}'

gcloud pubsub subscriptions list
````

Subscription delivery types: pull, push.
