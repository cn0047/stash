App Engine - Standard Environment Go 1.9
-

[doc](https://cloud.google.com/appengine/docs/standard/go/)
[config files](https://cloud.google.com/appengine/docs/standard/go/configuration-files)
[environments comparison](https://cloud.google.com/appengine/docs/the-appengine-environments)
[quotas](https://cloud.google.com/appengine/quotas)
[capabilities](https://cloud.google.com/appengine/docs/standard/go/capabilities/)
[pricing](https://cloud.google.com/appengine/pricing)
[requests limits](https://cloud.google.com/appengine/docs/standard/go/how-requests-are-handled#quotas_and_limits)
[services versions limits](https://cloud.google.com/appengine/docs/standard/python/an-overview-of-app-engine#limits)
[instance classes](https://cloud.google.com/appengine/docs/standard/#instance_classes)
[goland config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs)
[forum](https://groups.google.com/forum/#!forum/google-appengine-go)
[issue tracker](http://code.google.com/p/googleappengine/issues/list) [and](https://cloud.google.com/support/docs/issue-trackers)

````sh
# in web console
goapp serve app.yaml
````

````bash
# ~/.google-cloud-sdk/bin/dev_appserver.py \
#     --skip_sdk_update_check=false \
#     --log_level=debug \
#     --port=8080 --admin_port=8000 \
#     --storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
#     --default_gcs_bucket_name=itisgnp.appspot.com \
#     --go_debugging=true \
#     $(GOPATH)/src/go-app/.gae/app.yaml

service=cws-products
version=20181211t222125
instanceID=00c61b117c7f1d63b68246dded86378c0b3e12b6b46db84bd861492bc216e5d873e0408717

gcloud app describe

gcloud app services list

gcloud app instances list
gcloud app instances describe --service $service --version $version $instanceID
gcloud app instances ssh --service $service --version $version $instanceID

gcloud app versions list
gcloud app versions delete v1
# delete stopped versions:
for i in $( gcloud app versions list | grep STOPPED | awk '{print $2}' ); do
    gcloud app versions delete -q $i
done

gcloud app deploy
gcloud app deploy --verbosity=debug --project=thisismonitoring

gcloud app browse

gcloud app logs tail -s $service
````

## Overview

Warmup Requests - you can use to avoid latency while loading application code on a fresh instance.

Only internal appengine microservices have `X-Appengine-Inbound-Appid` header!

## +/-

Advantages:
* Pre-configured architecture.
* Automatically handle and balance all instances and data centers.
* Automatic scalability.
* Built-in web interface for task queues & crons.
* Pay for what you use.
* No server maintenance.
* Use CDN replication automatically (image, CSS, js content are replicated all across the globe).
* 1️⃣ App Engine SDK for locall development.
* 2️⃣ Free quota.

Disadvantages:
* http.DefaultTransport and http.DefaultClient are not available in App Engine
  (App Engine can only execute code called from an HTTP request).
* 60 second per request.
* templates - panic: open ../template/home.html: operation not permitted.
* impossible to write files, possible to open files for read only and only near main.go
  (open x.txt: no file writes permitted on App Engine).
* 1️⃣ `dev_appserver.py` - There are too many files in your application for changes in all of them to be monitored.
  You may have to restart the development server to see some changes to your files.
* 2️⃣ Free quota is tricky.
* No multithreading.
  [look here](https://blog.golang.org/go-and-google-app-engine)
  [and here](https://monosnap.com/file/Y66Cckm0pmQlG6GLEwoMxg684ig4RN)
* `(gcloud.app.instances.ssh)` Standard instances do not support this operation.
* No opportunity to debug gae-go.
* ERROR: (gcloud.app.deploy) INVALID_ARGUMENT: Your app may not have more than 15 versions. But recently was 20.
* ERROR: context deadline exceeded - from everywhere and it's annoying.
* Like the standard environment, the flexible environment does not support websockets.

Also, although goroutines and channels are present,
when a Go app runs on App Engine only one thread is run in a given instance.
That is, all goroutines run in a single operating system thread,
so there is no CPU parallelism available for a given client request.

## Memcache

[doc](https://cloud.google.com/appengine/docs/standard/go/memcache/)

````golang
item := memcache.Item{Key: key, Value: v, Expiration: expiration}
err := memcache.Set(ctx, &item)
````

## CRON

[doc](https://cloud.google.com/appengine/docs/standard/go/config/cron)
[deadlines](https://cloud.google.com/appengine/docs/standard/go/config/cronref#deadlines)

````yaml
cron:
- description: "daily summary job"
  url: /tasks/summary
  target: beta
  schedule: every 10 mins
  schedule: every 15 minutes
  schedule: every 24 hours
  schedule: every monday 09:00
  schedule: every 5 minutes from 10:00 to 14:00
  schedule: every day 08:00
````

## GO

````golang
"google.golang.org/appengine"
"google.golang.org/appengine/log"

ctx := appengine.NewContext(r)
module := appengine.ModuleName(ctx)
instance := appengine.InstanceID()

// r *http.Request
log.Infof(appengine.NewContext(r), "%#v", v)

appengine.AppID
appengine.DefaultVersionHostname 
appengine.IsDevAppServer()

err := runtime.RunInBackground(c, func(c appengine.Context) {
  // do something...
})

// delayed func
var expensiveFunc = delay.Func("some-arbitrary-key", func(ctx context.Context, a string, b int) {
        // do something expensive!
})
expensiveFunc.Call(ctx, "Hello, world!", 42)

// set something into context
c := context.WithValue(GAECtx, "key", "val")
c.Get("key")

ctx, _ := context.WithDeadline(cctx.GAECtx, time.Now().Add(time.Second*60))

// urlfetch
ctx := appengine.NewContext(r)
client := urlfetch.Client(ctx)
resp, err := client.Get(url)
````

Local Unit Testing:
````golang
# run
goapp test ./demos/transaction

"google.golang.org/appengine/aetest"

ctx, done, err := aetest.NewContext()

inst, err := aetest.NewInstance(nil)
defer inst.Close()
````
