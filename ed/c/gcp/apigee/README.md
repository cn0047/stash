Apigee
-

[docs](https://cloud.google.com/apigee/docs)
[docs](https://docs.apigee.com/api-platform/reference/apigee-reference)
[vars](https://cloud.google.com/apigee/docs/api-platform/reference/variables-reference)
[oauth](https://cloud.google.com/apigee/docs/api-platform/tutorials/secure-calls-your-api-through-oauth-20-client-credentials)
[examples](https://github.com/apigee/api-platform-samples)
[API](https://cloud.google.com/apigee/docs/reference/apis/apigee/rest)
[API operations](https://apidocs.apigee.com/operations)

Apigee - gateway.

Apigee environment deployment type: proxy, archive.

Organization - top-level container in Apigee.
It contains all API proxies and related resources.

Organization types: paid, evaluation.

KVM (Key Value Map) - not requires proxy redeployment.

````sh
gcloud apigee applications list
gcloud apigee deployments list
gcloud apigee environments list
gcloud apigee organizations list
gcloud apigee products list

# alpha

gcloud alpha apigee operations list

gcloud alpha apigee organizations list
gcloud alpha apigee organizations delete $org
gcloud alpha apigee organizations provision --authorized-network=default



# lint
npm i apigeelint
d=gh/ed/cloud/gcp/apigee/examples/API.One/apiproxy
apigeelint -s $d | jq
apigeelint --maxWarnings 0  --profile apigeex --formatter stylish.js --path $d

````

````sh
# vars
proxy.basepath      # /v1/weather
proxy.pathsuffix
client.scheme       # https
request.verb
request.header.host
request.uri
request.querystring
request.formparam.count

# property set
# where:
# $fnp - property file name without (.properties)
# $vn - var name from properties file
{propertyset.$fnp.$vn}

# tricks
{firstnonnull(target-path,proxy.pathsuffix)}

# js
httpClient.send(new Request(
'https://realtimelog.herokuapp.com:443/kaiuw6t7v8n', 'POST', {'Content-Type': 'application/json'},
JSON.stringify({code: 200, status: 'OK'})
));
````

````js
print(context.getVariable('foo'));
context.setVariable("response.header.X-Apigee-Target", context.getVariable("target.name"));
````

API key policy:
````
<APIKey ref="request.queryparam.apikey"/>
<APIKey ref="request.header.x-api-key"/>
````
