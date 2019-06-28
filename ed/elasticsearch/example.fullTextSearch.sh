# Full-text search

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : { "match" : { "about" : "rock climbing" } }
}'
# we'll receive: "I love to go rock climbing" and "I like to collect rock albums"

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    }
}'
# we'll receive: "I love to go rock climbing"

curl -XGET localhost:9200/megacorp/employee/_search -d '
{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    },
    "highlight": { "fields" : { "about" : {} } }
}
'
#

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "wildcard" : { "about" : "*limbing" }
    }
}'
# we'll receive: "I love to go rock climbing"

# Search by prefix
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "prefix" : { "about" : "REALLY" }
    }
}'
