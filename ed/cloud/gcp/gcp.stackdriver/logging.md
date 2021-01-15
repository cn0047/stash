Stackdriver Logging
-

[doc](https://cloud.google.com/logging/docs/quickstart-sdk)
[console](https://console.cloud.google.com/logs)
[quotas](https://cloud.google.com/logging/quotas)

````sh
gcloud logging logs list
gcloud logging metrics list
gcloud logging sinks list

gcloud logging write experiment "init."
gcloud logging write my-test-log "A simple entry"
gcloud logging write my-test-log-2 --payload-type=json \
  '{ "message": "My second entry", "weather": "partly cloudy"}'

gcloud logging read --limit=5
gcloud logging read --format=json --limit=50 --freshness=2d '
  resource.type=gae_app
  AND resource.labels.module_id="xws"
  AND logName="projects/xprod750/logs/appengine.googleapis.com%2Frequest_log"
  AND protoPayload.resource="/my-ep/part"
  AND timestamp>="2019-05-14T00:00:00.0Z" AND timestamp<="2019-05-15T00:00:00.0Z"
' | tee g.log | jq 'length'
````

Log view filter:
````sh
text:unicorn
text:unicorn text:phoenix          # unicorn or phoenix
text:"unicorn phoenix"             # unicorn and phoenix
text:uni*                          # wildcards: `?`, `*`
text:2017-02-05
text:200..299                      # `.` - like `.` in regex
text:unicorn text:NOT text:phoenix # unicorn NOT phoenix
text:unicorn text:OR text:phoenix  # unicorn OR phoenix
text:n=5                           # n=5

path:query             # `/query` or `/App/Query/17`
path:query unicorn     #  contain query and unicorn
path:query path:status # query or status

querystring:var=3
status:400..405

status:200
status:400..499 # HTTP status of 400 through 499
````

[Advanced filter](https://cloud.google.com/logging/docs/view/advanced-filters):
````sh
protoPayload.resource:"var=3"
protoPayload.status>=400 AND protoPayload.status <= 405
protoPayload.method="GET"

protoPayload.status=500

resource.type="gae_app"
resource.labels.module_id="cws-products"
logName="projects/itismonitoring/logs/appengine.googleapis.com%2Frequest_log"
protoPayload.resource!="/health-check"
protoPayload.resource!="/cron/addPingJobs"
protoPayload.resource!="/worker/ping"
protoPayload.resource!="/static/"
severity>=CRITICAL
severity>=ERROR
severity>=WARNING
NOT "GET /sitemap/"
````

Go:
````golang
"cloud.google.com/go/logging"

l, err := logging.NewClient(ctx, "thisismonitoring")
if err != nil {
    log.Errorf(ctx, "Failed to create log client: %v", err)
    return
}
defer l.Close()
logger := l.Logger("namespace").StandardLogger(logging.Info)
logger.Printf("[💻] ✅ Visit # %v.", 1)
````
