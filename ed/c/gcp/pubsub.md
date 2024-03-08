Pub/Sub
-

[docs](https://cloud.google.com/pubsub/docs/overview)
[SLA](https://cloud.google.com/pubsub/sla)

````sh
# emulator
gcloud components install pubsub-emulator
gcloud beta emulators pubsub start
gcloud beta emulators pubsub env-init

# start docker
docker run --rm -ti -p 8681:8681 --name pubsub-emulator \
  -e PUBSUB_PROJECT1="test-project,test-topic:test-subscription" \
  messagebird/gcloud-pubsub-emulator:latest

# stop
docker stop pubsub-emulator

export PUBSUB_EMULATOR_HOST=localhost:8085
export PUBSUB_EMULATOR_HOST=localhost:8681
#
gcloud config set auth/disable_credentials true
gcloud config set api_endpoint_overrides/pubsub "http://localhost:8085/"
gcloud config set api_endpoint_overrides/pubsub "http://localhost:8681/"



topic=test
gcloud pubsub topics create $topic
gcloud pubsub topics list

gcloud pubsub topics publish $topic --message='{"type":"test","data":"ok"}'


gcloud pubsub subscriptions create --topic=$topic $subscription
gcloud pubsub subscriptions list

gcloud pubsub subscriptions pull --auto-ack --limit=1 $subscription
````

Pub/sub topic - doesn't have region configuration (it's global beneath).
Pub/sub subscription - doesn't have region configuration as well.

Subscription delivery types: pull, push, write to BigQuery, write to Cloud Storage.
Subscription filter: `attributes.key="foo"`.
Subscription has: exactly once delivery.
Subscription has: dead lettering.
Subscription has: retry after exponential backoff delay.
