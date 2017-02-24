Reindex (with ALIAS) wihout downtime
-

export host='localhost'
export port=9201
export index=ziipr
export alias=ziipr
export type=users

````json
curl http://$host:$port/_cat/indices?v

curl -XPUT http://$host:$port/$alias/ -d '{
  "mappings" : {
    "users": {
      "properties": {
        "user_id": {"type": "long"},
        "email": {"type": "string", "index": "not_analyzed"},
        "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
        "updated_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
        "last_seen_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
        "location": {"type": "geo_point", "lat_lon": "true"},
        "profile_meta_data": {"type": "nested"},
        "photos": {"type": "nested"},
        "video": {
          "type": "nested",
          "properties": {
            "id": {"type": "string"},
            "type_id": {"type": "string"},
            "status": {"type": "string"},
            "file_name": {"type": "string"},
            "sproutVideo": {"type": "nested"}
          }
        }
      }
    },
    "events": {
      "properties": {
        "user_id": {"type": "long"},
        "type_id": {"type": "integer"},
        "year": {"type": "integer"},
        "month": {"type": "integer"},
        "day": {"type": "integer"},
        "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"}
      }
    }
  }
}'

curl -XGET $host:$port/$index/$type/_count

php cli.php elasticSearch reIndexBulk users $alias users
php cli.php elasticSearch reIndexBulk events ziipr_v2 events

curl -XGET $host:$port/$alias/$type/_count -d '{"query": {"match_all" : {}}}'

curl -XDELETE http://$host:$port/ziipr/

curl -XPOST localhost:$port/_aliases -d '{
"actions": [
    {"add": {"alias": "ziipr", "index": "ziipr_v2"}}
]
}'

curl -XGET $host:$port/_alias/
curl -XGET $host:$port/_cat/aliases?v
````



## reindex (with ALIAS) wihout downtime

curl http://internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/_cat/indices?v

curl -XPUT http://internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/ziipr_v2/ -d '{
"mappings" : {
    "users": {
        "properties": {
            "user_id": {"type": "long"},
            "email": {"type": "string", "index": "not_analyzed"},
            "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
            "updated_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
            "last_seen_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"},
            "location": {"type": "geo_point", "lat_lon": "true"},
            "profile_meta_data": {"type": "nested"},
            "photos": {"type": "nested"},
            "video": {
                "type": "nested",
                "properties": {
                    "id": {"type": "string"},
                    "type_id": {"type": "string"},
                    "status": {"type": "string"},
                    "file_name": {"type": "string"},
                    "sproutVideo": {"type": "nested", "null_value": "{}"}
                }
            }
        }
    },
    "events": {
        "properties": {
            "user_id": {"type": "long"},
            "type_id": {"type": "integer"},
            "year": {"type": "integer"},
            "month": {"type": "integer"},
            "day": {"type": "integer"},
            "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"}
        }
    }
}
}'

curl -XGET internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/ziipr/$type/_count -d '{"query": {"match_all" : {}}}'

php cli.php elasticSearch reIndexBulk users ziipr_v2 users
php cli.php elasticSearch reIndexBulk events ziipr_v2 events

curl -XGET internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/ziipr_v2/$type/_count -d '{"query": {"match_all" : {}}}'

curl -XDELETE http://internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/ziipr/

curl -XPOST internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/_aliases -d '{
"actions": [
    {"add": {"alias": "ziipr", "index": "ziipr_v2"}}
]
}'

curl -XGET -i internal-prod-ziipr-elasticsearch-lb-798730478.eu-west-1.elb.amazonaws.com:9201/_alias/
