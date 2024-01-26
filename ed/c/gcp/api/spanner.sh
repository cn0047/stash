# Spanner API

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



# list backups
curl -X GET -H $AUTH "$h/v1/projects/$prj/instances/$i/backups" | jq

# create backup
id='test-backup-1'
dbn="projects/$prj/instances/$i/databases/$db"
curl -X POST -H $AUTH -H $jh "$h/v1/projects/$prj/instances/$i/backups?backupId="$id -d '{
  "database": "'$dbn'",
  "expireTime": "2023-01-01T00:00:00.1Z"
}'
