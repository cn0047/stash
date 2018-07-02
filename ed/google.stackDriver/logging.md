Stackdriver Logging
-

[doc](https://cloud.google.com/logging/docs/quickstart-sdk)
[console](https://console.cloud.google.com/logs)

````sh
gcloud logging write my-test-log "A simple entry"
gcloud logging write my-test-log-2 --payload-type=json \
  '{ "message": "My second entry", "weather": "partly cloudy"}'
````

````go
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
