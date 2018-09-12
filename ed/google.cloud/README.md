Google Cloud
-

[doc](https://cloud.google.com/sdk/gcloud/)

````bash
# ~/.google-cloud-sdk/bin/gcloud

# login
gcloud auth login
gcloud auth application-default login
gcloud auth list
# or
gcloud auth activate-service-account --key-file=account.json

gcloud config list
gcloud config set project thisismonitoring
gcloud config set disable_prompts false
gcloud config set disable_usage_reporting true

gcloud projects list --format="json"
gcloud projects list \
  --format="table(projectNumber,projectId,createTime)" \
  --filter="createTime.date('%Y-%m-%d', Z)='2016-05-11'"

gcloud components list
gcloud components update
gcloud components install app-engine-php

gcloud source repos list

gcloud container clusters list

gcloud compute instances list
````
