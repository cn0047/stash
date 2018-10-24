Google Cloud
-

[doc](https://cloud.google.com/sdk/gcloud/)

````bash
# ~/.google-cloud-sdk/bin/gcloud

export GOOGLE_APPLICATION_CREDENTIALS=${CURDIR}/serviceAccount.json

~/.google-cloud-sdk/bin/dev_appserver.py \
  --skip_sdk_update_check=false \
  --log_level=debug \
  --port=8080 --admin_port=8000 \
  --storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
  --default_gcs_bucket_name=itisgnp.appspot.com \
  --go_debugging=true \
  $(GOPATH)/src/go-app/app.yaml

gcloud components list
gcloud components update
gcloud components install app-engine-php

gcloud source repos list

gcloud container clusters list

gcloud compute instances list

# login
gcloud auth login
gcloud auth application-default login
gcloud auth activate-service-account --key-file=service-account.json
gcloud auth list

gcloud config configurations list
gcloud config list
gcloud config set project thisismonitoring
gcloud config set disable_prompts false
gcloud config set disable_usage_reporting true

gcloud projects list --format="json"
gcloud projects list \
  --format="table(projectNumber,projectId,createTime)" \
  --filter="createTime.date('%Y-%m-%d', Z)='2016-05-11'"

gcloud beta projects get-iam-policy itisgnp
````
