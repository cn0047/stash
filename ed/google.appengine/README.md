App Engine
-

[logs](https://console.cloud.google.com/logs)
[GoLand config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs)

````
# ~/.google-cloud-sdk/bin/gcloud

# login
gcloud auth login
gcloud auth list
# or
gcloud auth activate-service-account --key-file=account.json

gcloud config list
gcloud config set project thisismonitoring

gcloud projects list

gcloud components list
gcloud components update
gcloud components install app-engine-php

gcloud source repos list

gcloud app instances list
gcloud app services list

gcloud app deploy
gcloud app deploy --verbosity=debug --project=thisismonitoring

gcloud app browse

gcloud app logs tail -s default
````

````
# in web console
goapp serve app.yaml
````

Warmup Requests - you can use to avoid latency while loading application code on a fresh instance.

Only internal appengine microservices have `X-Appengine-Inbound-Appid` header!


## Google Cloud SQL

MySQL.

## Google Cloud Storage (Cloud file storage)
