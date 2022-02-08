Apigee
-

[docs](https://cloud.google.com/apigee/docs)

Apigee - gateway.

````sh
gcloud apigee applications list
gcloud apigee deployments list
gcloud apigee environments list
gcloud apigee organizations list
gcloud apigee products list

gcloud alpha apigee operations list

````

````sh
# curl internal apigee proxy

PROJECT_ID='' # !!!
SA='' # !!!
SUBNET='default'
PROJECT_NUMBER=$(gcloud projects describe $PROJECT_ID --format="value(projectNumber)")
TOKEN=$(gcloud auth print-access-token)
AUTH="Authorization: Bearer $TOKEN"
# get apigee instance
curl -H $AUTH "https://apigee.googleapis.com/v1/organizations/$PROJECT_ID/instances"

# use name from output
INSTANCE_NAME='us-central1'

# find zone for location from output
gcloud compute zones list
ZONE='us-central1-a'

# create vm
gcloud beta compute --project=$PROJECT_ID instances create $INSTANCE_NAME \
  --zone=$ZONE \
  --machine-type=e2-micro \
  --subnet=$SUBNET \
  --network-tier=PREMIUM \
  --no-restart-on-failure \
  --maintenance-policy=TERMINATE \
  --preemptible \
  --service-account=$SA \
  --scopes=https://www.googleapis.com/auth/cloud-platform \
  --tags=http-server,https-server \
  --image=debian-10-buster-v20210217 \
  --image-project=debian-cloud \
  --boot-disk-size=10GB \
  --boot-disk-type=pd-standard \
  --boot-disk-device-name=$INSTANCE_NAME \
  --no-shielded-secure-boot \
  --shielded-vtpm \
  --shielded-integrity-monitoring \
  --reservation-affinity=any

# delete if needed
gcloud beta compute --project=$PROJECT_ID instances delete $INSTANCE_NAME

# ssh
gcloud compute ssh $INSTANCE_NAME --zone=$ZONE --project=$PROJECT_ID

# in ssh
sudo apt-get update -y
sudo apt-get install -y jq
PROJECT_ID='' # !!!
TOKEN=$(gcloud auth print-access-token)
AUTH="Authorization: Bearer $TOKEN"
ENV_GROUP_HOSTNAME=$(curl -H "$AUTH" "https://apigee.googleapis.com/v1/organizations/$PROJECT_ID/envgroups" -s | jq -r '.environmentGroups[0].hostnames[0]')
INTERNAL_LOAD_BALANCER_IP=$(curl -H "$AUTH" "https://apigee.googleapis.com/v1/organizations/$PROJECT_ID/instances" -s | jq -r '.instances[0].host')

# download certificate
curl -H "$AUTH" "https://apigee.googleapis.com/v1/organizations/$PROJECT_ID" | jq -r .caCertificate | base64 -d > cacert.crt

# curl
curl -H "Host: $ENV_GROUP_HOSTNAME" "https://example.$PROJECT_ID.apigee.internal/hw" \
  --resolve "example.$PROJECT_ID.apigee.internal:443:$INTERNAL_LOAD_BALANCER_IP" \
  --cacert cacert.crt

````
