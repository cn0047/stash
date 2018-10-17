Cloud Storage
-

````sh
# ls /Users/k/.google-cloud-sdk/bin/gsutil

gsutil acl

gsutil ls
gsutil ls gs://my-awesome-bucket
gsutil ls -l gs://my-awesome-bucket/kitten.png
gsutil rm gs://my-awesome-bucket/kitten.png

gsutil rsync -d -r s3://my-aws-bucket gs://example-bucket
````
