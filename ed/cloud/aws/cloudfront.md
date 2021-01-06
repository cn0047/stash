CloudFront (CDN)
-

````sh
aws cloudfront list-distributions

aws cloudfront get-distribution --id $id
````

HTTP Response headers:
````sh
x-cache: Miss from cloudfront
via: 1.1 9b097dfab92228268a37145aac5629c1.cloudfront.net (CloudFront)

x-cache: Hit from cloudfront
via: 1.1 7b32163caf7e91fe96df7bbeaa58c0f9.cloudfront.net (CloudFront)
````
