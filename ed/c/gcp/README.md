Google Cloud
-

[doc](https://cloud.google.com/sdk/gcloud/)
[code samples](https://cloud.google.com/docs/samples)
[products](https://cloud.google.com/products/)
[console](https://console.cloud.google.com/)
[API clients](https://developers.google.com/api-client-library/)
[go libraries](https://pkg.go.dev/cloud.google.com/go#section-readme)
[go libraries](https://github.com/googleapis)
[locations](https://cloud.google.com/about/locations/)
[machine types](https://cloud.google.com/compute/docs/machine-types)
[tags by image](https://cloud.google.com/vision/docs/drag-and-drop)
[measure latency](https://gcping.com/)
[certification](https://cloud.google.com/certification/cloud-developer)
[resource manager](https://console.cloud.google.com/cloud-resource-manager)

<img src="https://gist.github.com/cn007b/384d6938ebef985347b29c15476b55c5/raw/7b19f9797d647c75fdbd7098a359d8788e0a7107/gcp.funcOrRun.png" width="70%" />

````bash
# in URL:
# ?project=PROJECT_ID&authuser=2

# ~/.google-cloud-sdk/bin/gcloud

echo $GOOGLE_APPLICATION_CREDENTIALS
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

gcloud compute zones list
gcloud compute regions list

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
gcloud auth print-access-token
gcloud auth print-identity-token
gcloud auth print-identity-token --impersonate-service-account=$sa --audiences=$cloudRunURL --include-email
gcloud auth revoke --all # logout
gcloud auth application-default revoke # logout

# default service account
ls -la ~/.config/gcloud/application_default_credentials.json

gcloud config configurations create my_cfg
gcloud config configurations activate my_cfg
gcloud auth login
gcloud auth application-default login
gcloud config configurations list
gcloud config configurations delete my_cfg
gcloud config list
gcloud config set project itismonitoring
gcloud config set project thisisrealtimelog
gcloud config set disable_prompts false
gcloud config set disable_usage_reporting true
# ls configurations
ls ~/.config/gcloud/configurations/

export PROJECT_ID=`gcloud config get project`

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
* GAE - App Engine ([PaaS](https://gist.githubusercontent.com/cn0047/384d6938ebef985347b29c15476b55c5/raw/a1b334180b6aaa745c22626a9d3489f840569edd/cloudComputingTypes.jpeg))
* GKE - Container Engine (PaaS)
* GCE - Compute Engine (IaaS)
* Cloud Functions (Serverless)

.gcloudignore - holds files to be ignored during uploading directory to GCP.

Managed instances groups used for LB.

# Networking

34 regions now.
103 zones now.

Google Cloud Interconnect extends your on-premises network
to Google's network through a highly available, low latency connection.

Each VPC network is a global entity spanning all GCP regions.
Each VPC network is subdivided into subnets,
and each subnet is contained within a single region.
Each subnet has a contiguous private RFC1918 IP space.
Virtual machine instances in a VPC network can communicate with instances
in all other subnets of the same VPC network, regardless of region,
using their RFC1918 private IP addresses.
You can isolate portions of the network, even entire subnets, using firewall rules.

Max 5 network per GCP project.
Internal Ip address used internal DNS server for FQDN.
To scale Cloud VPN - add multiple tunnels.
Subnets can not span across multiple regions.
Cloud Router doesn't use BGP link to advertise network changes.
Google Virtual Network Subnets are regional resource.

# ESP (Extensible Service Proxy)

ESP - Envoy-based high-performance, scalable proxy that runs in front of
OpenAPI or gRPC API backend and provides API management features.

# NAT (Network Address Translation)
