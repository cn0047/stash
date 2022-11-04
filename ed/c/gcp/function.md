Cloud Function
-

[pricing](https://cloud.google.com/functions/pricing)

````sh
# prepare deploy
f=ed/l/go/examples/3rdparty/gcp/function/hw.go
source=/tmp/x/
mkdir -p $source && cp $f $source

# deploy
# source: dir or .zip or source repository
gcloud functions deploy hw \
  --runtime=go116 \
  --trigger-http \
  --allow-unauthenticated \
  --source=$source

gcloud functions list
gcloud functions describe hw
````

Function is per region.

Trigger type: HTTP, PubSub, GCS, Firestore, Firebase.

HTTP trigger looks like: `https://$region-$prj.cloudfunctions.net/$funcName`.
