APIs
-

[docs](https://cloud.google.com/apis/docs/getting-started)
[design](https://cloud.google.com/apis/design)
[library](https://console.cloud.google.com/apis/library/browse)
[dashboard](https://console.cloud.google.com/apis/dashboard)

````sh
PROJECT_ID=''
prj=$PROJECT_ID

# get token
TOKEN=$(gcloud auth print-access-token)
AUTH="Authorization: Bearer $TOKEN"

jh='content-type: application/json'

h="https://www.googleapis.com"

# list NEGs
region=us-central1
curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/regions/$region/networkEndpointGroups" | jq

curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/global/sslCertificates" | jq

curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/global/backendServices" | jq

curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/global/urlMaps" | jq

curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/global/targetHttpsProxies" | jq

curl -X GET -H $AUTH "$h/compute/beta/projects/$prj/global/forwardingRules" | jq

````
