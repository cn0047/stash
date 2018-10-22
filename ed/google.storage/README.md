Cloud Storage
-

[docs appengine](https://cloud.google.com/appengine/docs/standard/go/googlecloudstorageclient/read-write-to-cloud-storage)
[permissions](https://cloud.google.com/storage/docs/access-control/iam-permissions)

````sh
# ls /Users/k/.google-cloud-sdk/bin/gsutil

gsutil acl get gs://itisgnp.appspot.com
gsutil acl ch -u itisgnp@appspot.gserviceaccount.com:O gs://itisgnp.appspot.com

gsutil iam get gs://itisgnp.appspot.com
# gsutil iam ch serviceAccount:itisgnp@appspot.gserviceaccount.com:objectCreator gs://itisgnp.appspot.com
# gsutil iam ch user:cn007b@gmail.com:objectCreator gs://itisgnp.appspot.com

gsutil ls gs://itisgnp.appspot.com
gsutil ls -l gs://itisgnp.appspot.com/test
gsutil rm gs://itisgnp.appspot.com/test/none

gsutil rsync -d -r s3://my-aws-bucket gs://example-bucket
````
