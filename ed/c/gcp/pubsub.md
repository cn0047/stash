Pub/Sub
-

````sh
# emulator
gcloud components install pubsub-emulator
#
gcloud beta emulators pubsub start
gcloud beta emulators pubsub start --project=test-project
#
export PUBSUB_EMULATOR_HOST=localhost:8085



gcloud pubsub topics list
gcloud pubsub topics publish $topic --message='{"type":"test","data":"ok"}'

gcloud pubsub subscriptions list

gcloud pubsub subscriptions pull --auto-ack --limit=1 $subscription
````

Pub/sub topic - doesn't have region configuration (it's global beneath).
Pub/sub subscription - doesn't have region configuration as well.

Subscription delivery types: pull, push.
Subscription has filter (`attributes.key="foo"`).
Subscription has retry after exponential backoff delay.
