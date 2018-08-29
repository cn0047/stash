App Engine
-

[doc](https://cloud.google.com/appengine/docs/standard/go/)
[console](https://console.cloud.google.com/)
[environments comparison](https://cloud.google.com/appengine/docs/the-appengine-environments)
[quotas](https://cloud.google.com/appengine/quotas)
[capabilities](https://cloud.google.com/appengine/docs/standard/go/capabilities/)
[pricing](https://cloud.google.com/appengine/pricing)
[requests limits](https://cloud.google.com/appengine/docs/standard/go/how-requests-are-handled#quotas_and_limits)
[services versions limits](https://cloud.google.com/appengine/docs/standard/python/an-overview-of-app-engine#limits)
[instance classes](https://cloud.google.com/appengine/docs/standard/#instance_classes)
[machine types](https://cloud.google.com/compute/docs/machine-types)
[config files](https://cloud.google.com/appengine/docs/flexible/go/reference/app-yaml)
[goland config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs)
[forum](https://groups.google.com/forum/#!forum/google-appengine-go)
[issue tracker](http://code.google.com/p/googleappengine/issues/list)

````
# in web console
goapp serve app.yaml
````

````bash
# ~/.google-cloud-sdk/bin/dev_appserver.py \
#     --skip_sdk_update_check=false \
#     --log_level=debug \
#     --port=8080 --admin_port=8000 \
#     --storage_path=$(GOPATH)/.data --support_datastore_emulator=false \
#     --go_debugging=true \
#     $(GOPATH)/src/go-app/.gae/app.yaml

gcloud app describe

gcloud app instances list
gcloud app services list

gcloud app versions list
gcloud app versions delete v1
# delete stopped versions:
for i in $( gcloud app versions list | grep STOPPED | awk '{print $2}' ); do
    gcloud app versions delete -q $i
done

gcloud app deploy
gcloud app deploy --verbosity=debug --project=thisismonitoring

gcloud app browse

gcloud app logs tail -s default

gcloud app instances list
instanceID=00c61b117c2d2d8a7c1c04ec77e3f133cecb6a72ecdf1da0844b8223f622b9d72227fe
gcloud app instances ssh --service default --version 20180727t232426 $instanceID
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
* 1️⃣ `dev_appserver.py` - There are too many files in your application for changes in all of them to be monitored.
  You may have to restart the development server to see some changes to your files.
* 2️⃣ Free quota is tricky.
* No multithreading?
  [look here](https://blog.golang.org/go-and-google-app-engine)
  [and here](https://monosnap.com/file/Y66Cckm0pmQlG6GLEwoMxg684ig4RN)
* `(gcloud.app.instances.ssh)` Standard instances do not support this operation.
* No opportunity to debug gae-go.

Also, although goroutines and channels are present,
when a Go app runs on App Engine only one thread is run in a given instance.
That is, all goroutines run in a single operating system thread,
so there is no CPU parallelism available for a given client request.

## Memcache

Avoid Memcache hot keys.
Hot keys are a common anti-pattern that can cause Memcache capacity to be exceeded.

For Dedicated Memcache, we recommend that the peak access rate on a single key
should be 1-2 orders of magnitude less than the per-GB rating.
For example, the rating for 1 KB sized items is 10,000 operations per second per GB of Dedicated Memcache.
Therefore, the load on a single key should not be higher
than 100 - 1,000 operations per second for items that are 1 KB in size.

## GO

````go
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
````

Local Unit Testing:

````go
# run
goapp test ./demos/transaction

"google.golang.org/appengine/aetest"

ctx, done, err := aetest.NewContext()

inst, err := aetest.NewInstance(nil)
defer inst.Close()
````
