# Full-text search

url=localhost:9200/megacorp/employee

curl -XGET $url/_search -d '{
    "query" : { "match" : { "about" : "rock climbing" } }
}'
# result: "I love to go rock climbing" and "I like to collect rock albums"

curl -XGET $url/_search -d '{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    }
}'
# result: "I love to go rock climbing"

curl -XGET $url/_search -d '
{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    },
    "highlight": { "fields" : { "about" : {} } }
}
'

curl -XGET $url/_search -d '{
    "query" : {
        "wildcard" : { "about" : "*limbing" }
    }
}'
# result: "I love to go rock climbing"

# Search by prefix
curl -XGET $url/_search -d '{
    "query" : {
        "prefix" : { "about" : "REALLY" }
    }
}'
