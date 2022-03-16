Apigee
-

[docs](https://cloud.google.com/apigee/docs)
[docs](https://docs.apigee.com/api-platform/reference/apigee-reference)
[oauth](https://cloud.google.com/apigee/docs/api-platform/tutorials/secure-calls-your-api-through-oauth-20-client-credentials)
[examples](https://github.com/apigee/api-platform-samples)
[API](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest)
[API operations](https://apidocs.apigee.com/operations)

Apigee - gateway.

Apigee environment deployment type: proxy, archive.

````sh
gcloud apigee applications list
gcloud apigee deployments list
gcloud apigee environments list
gcloud apigee organizations list
gcloud apigee products list

gcloud alpha apigee operations list

````

API key policy:
````
<APIKey ref="request.queryparam.apikey"/>
<APIKey ref="request.header.x-api-key"/>
````