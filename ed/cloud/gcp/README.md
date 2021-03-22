Google Cloud
-

[doc](https://cloud.google.com/sdk/gcloud/)
[products](https://cloud.google.com/products/)
[console](https://console.cloud.google.com/)
[APIs](https://developers.google.com/api-client-library/)
[locations](https://cloud.google.com/about/locations/)
[machine types](https://cloud.google.com/compute/docs/machine-types)
[tags by image](https://cloud.google.com/vision/docs/drag-and-drop)

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

gcloud app describe

gcloud components list
gcloud components update
gcloud components install app-engine-php

gcloud services list
gcloud services enable cloudprofiler.googleapis.com

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
gcloud config set project itismonitoring
gcloud config set project thisisrealtimelog
gcloud config set disable_prompts false
gcloud config set disable_usage_reporting true

gcloud projects list --format="json"
gcloud projects list \
  --format="table(projectNumber,projectId,createTime)" \
  --filter="createTime.date('%Y-%m-%d', Z)='2016-05-11'"
gcloud projects describe cl-dev

gcloud beta projects get-iam-policy itisgnp
````

````
Organization ⇒ Folder ⇒ Project ⇒ Resource
````

Computing service:
* GAE - App Engine ([PaaS](https://twitter.com/cn007b/status/1024010042838851585))
* GKE - Container Engine (PaaS)
* GCE - Compute Engine (IaaS)
* Cloud Functions (Serverless)

#### LB (Load Balancer)

Global external LB: HTTP(S), SSL, TCP Proxy.
Regional external LB: Network, Internal.

* HTTP(S)
* TCP
* UDP

Managed instances groups used for LB.

#### Networking

Google Cloud Interconnect extends your on-premises network
to Google's network through a highly available, low latency connection.

Max 5 network per GCP project.
Internal Ip address used internal DNS server for FQDN.
To scale Cloud VPN - add multiple tunnels.
Subnets can not span across multiple regions.
Cloud Router doesn't use BGP link to advertise network changes.
Google Virtual Network Subnets are regional resource.
