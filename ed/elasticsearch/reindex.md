Reindex (with ALIAS) wihout downtime
-

````json
curl http://localhost:9200/_cat/indices?v

curl -XPUT http://localhost:9200/ziipr_v2/ -d '{
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
            "photos": {"type": "nested"}
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

curl -XGET localhost:9200/ziipr/users/_count -d '{"query": {"match_all" : {}}}'

php cli.php elasticSearch reIndexBulk users ziipr_v2 users
php cli.php elasticSearch reIndexBulk events ziipr_v2 events

curl -XGET localhost:9200/ziipr_v2/users/_count -d '{"query": {"match_all" : {}}}'

curl -XDELETE http://localhost:9200/ziipr/

curl -XPOST localhost:9200/_aliases -d '{
"actions": [
    {"add": {"alias": "ziipr", "index": "ziipr_v2"}}
]
}'

curl -XGET localhost:9200/_alias/
curl -XGET localhost:9200/_cat/aliases?v
````
