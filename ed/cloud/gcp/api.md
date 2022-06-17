APIs
-

[docs](https://cloud.google.com/apis/docs/getting-started)
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

## Spanner

````sh
h="https://spanner.googleapis.com"



# SESSION
# create
db='testdb'
sess=`curl -X POST -H $AUTH "$h/v1/projects/$prj/instances/$i/databases/$db/sessions" | jq -r '.name'`
echo $sess
# list
curl -X POST -H $AUTH "$h/v1/projects/$prj/instances/$i/databases/$db/sessions"
# get
curl -X GET -H $AUTH "$h/v1/projects/$prj/instances/$i/databases/$db/sessions"
# delete
curl -X DELETE -H $AUTH "$h/v1/${sess}"



# list instances
curl -X GET -H $AUTH "$h/v1/projects/$prj/instances" | jq

# list databases
i='main'
curl -X GET -H $AUTH "$h/v1/projects/$prj/instances/$i/databases" | jq

# query read
curl -X POST -H $AUTH -H $jh "$h/v1/${sess}:executeSql" -d '{
  "sql": "SELECT * FROM test",
  "params": {},
  "paramTypes": {}
}' | jq

# query read with params
curl -X POST -H $AUTH -H $jh "$h/v1/${sess}:executeSql" -d '{
  "sql": "SELECT * FROM test WHERE id = @id",
  "params": {"id": "1"},
  "paramTypes": {"id": {"code": "INT64"}}
}' | jq

# query write
# start transaction
curl -X POST -H $AUTH -H $jh "$h/v1/${sess}:beginTransaction" -d '{
  "options": {"readWrite": {}}
}'
# commit
t="QVAxb2RaX0FQMEdBRHM2X0pCVEJtVlh5UTJPMy1mUFRnWHY2R0E="
curl -X POST -H $AUTH -H $jh "$h/v1/${sess}:commit" -d '{
  "mutations": [
    {
      "insert": {
        "table": "test",
        "columns": ["id", "msg"],
        "values": [["2", "cli"]]
      }
    }
  ],
  "transactionId": "'$t'"
}'
````
