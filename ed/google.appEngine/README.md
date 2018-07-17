App Engine
-

[doc](https://cloud.google.com/appengine/docs/standard/go/)
[console](https://console.cloud.google.com/)
[quotas](https://cloud.google.com/appengine/quotas)
[GoLand config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs)

````
# ~/.google-cloud-sdk/bin/gcloud

# login
gcloud auth login
gcloud auth list
# or
gcloud auth activate-service-account --key-file=account.json

gcloud config list
gcloud config set project thisismonitoring

gcloud projects list

gcloud components list
gcloud components update
gcloud components install app-engine-php

gcloud source repos list

gcloud app instances list
gcloud app services list

gcloud app versions list
gcloud app versions delete v1

gcloud app deploy
gcloud app deploy --verbosity=debug --project=thisismonitoring

gcloud app browse

gcloud app logs tail -s default
````

````
gcloud compute instances list
````

````
# in web console
goapp serve app.yaml
````

Warmup Requests - you can use to avoid latency while loading application code on a fresh instance.

Only internal appengine microservices have `X-Appengine-Inbound-Appid` header!

## +/-

Advantages:

Disadvantages:
* http.DefaultTransport and http.DefaultClient are not available in App Engine.
* There are too many files in your application for changes in all of them to be monitored.
  You may have to restart the development server to see some changes to your files.

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

````go

````
