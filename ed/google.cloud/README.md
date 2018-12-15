Google Cloud
-

[doc](https://cloud.google.com/sdk/gcloud/)
[products](https://cloud.google.com/products/)
[console](https://console.cloud.google.com/)
[APIs](https://developers.google.com/api-client-library/)
[locations](https://cloud.google.com/about/locations/)
[machine types](https://cloud.google.com/compute/docs/machine-types)

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
gcloud projects describe clique-dev

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

## GCE (Compute Engine)

* High performance virtual machine.
* Persistent disc storage.
* Local SSD & RAM.
* Supports glabal balancing.
* Supports auto scaling.
* Supports CDN.
* Auto restart.

[Machine types](https://cloud.google.com/compute/docs/machine-types):
Shared core machines: f1-micro, g1-small.

Predefined Virtual Machines types:
* standard
* shared-core
* high-cpu
* high-memory

It's possible to attach 10 per CPU persistent disks to GCE VM.

Instance template - define machine type, image, zone
and other properties.

Instance groups:
* managed (zonal, regional)
* unmanaged

#### LB (Load Balancer)

Global external LB: HTTP(S), SSL, TCP Proxy.
Regional external LB: Network, Internal.

* HTTP(S)
* TCP
* UDP

Managed instances groups used for LB.

#### CKE (Container Engine)

#### Networking

BGP       - Border Gateway Protocol.
Cloud DNS - Domain Name System.
ECMP      - Equal - Cost Multi - Path routing.
FQDN      - Fully Qualified Domain Name.
MPLS      - Multiprotocol Label Switching.
NAT       - Network Address Translation.
VPC       - Virtual Private Cloud.
VPS       - Virtual Private Server.

Google Cloud Interconnect extends your on-premises network
to Google's network through a highly available, low latency connection.

Max 5 network per GCP project.
Internal Ip address used internal DNS server for FQDN.
To scale Cloud VPN - add multiple tunnels.
Subnets can not span across multiple regions.
Cloud Router doesn't use BGP link to advertise network changes.
Cloud VPN is not regional Service.
Google Virtual Network Subnets are regional resource.
Cloud VPN is regional Service.
