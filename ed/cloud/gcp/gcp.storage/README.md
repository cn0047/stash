Cloud Storage
-

[docs](https://cloud.google.com/storage/docs/)
[pricing](https://cloud.google.com/storage/pricing)
[quotas](https://cloud.google.com/storage/quotas)
[appengine](https://cloud.google.com/appengine/docs/standard/go/googlecloudstorageclient/read-write-to-cloud-storage)
[permissions](https://cloud.google.com/storage/docs/access-control/iam-permissions)
[mount](https://cloud.google.com/storage/docs/gcs-fuse)
[object versioning](https://cloud.google.com/storage/docs/using-object-versioning)
[wildcards](https://cloud.google.com/storage/docs/gsutil/addlhelp/WildcardNames)

Cloud Storage - like AWS S3.

````sh
# ls /Users/k/.google-cloud-sdk/bin/gsutil

gsutil acl get gs://itisgnp.appspot.com
gsutil acl ch -u itisgnp@appspot.gserviceaccount.com:O gs://itisgnp.appspot.com

gsutil iam get gs://itisgnp.appspot.com
# gsutil iam ch serviceAccount:itisgnp@appspot.gserviceaccount.com:objectCreator gs://itisgnp.appspot.com
# gsutil iam ch user:cn007b@gmail.com:objectCreator gs://itisgnp.appspot.com

gsutil ls gs://itisgnp.appspot.com
gsutil ls -l gs://itisgnp.appspot.com/test
gsutil ls gs://example-bucket/** | wc -l

gsutil rm gs://itisgnp.appspot.com/test/none

gsutil du -sh gs://example-bucket
gsutil hash local-file
gsutil stat gs://example-bucket/composite-object

gsutil -m cp -R top-level-dir gs://example-bucket
gsutil -m cp -R top-level-dir/subdir/image* gs://example-bucket

gsutil -m rsync -r local-dir gs://example-bucket
gsutil rsync -d -r s3://my-aws-bucket gs://example-bucket

gsutil notification list gs://[BUCKET_NAME]
gsutil notification create -t [TOPIC_NAME] -f json gs://[BUCKET_NAME]
````

Transcoding of gzip-compressed files:
````
Content-Type: text/plain
Content-Encoding: gzip
````

Object max size limit - 5 TB.

Storage class:
* multi-regional
* regional
* nearline (infrequently accessed data, backup)
* coldline (rarely accessed data, disaster recovery)

Frequency limit to create/delete bucket is 2 second.
Frequency limit to update object is 1 second.
For nearline min storage duration - 30 days.

Cloud spanner offers ACID transactions and can scale to thousands of nodes.

## +/-

Advantages:
* versioning
