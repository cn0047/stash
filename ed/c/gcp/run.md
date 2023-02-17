Cloud Run
-

[docs](https://cloud.google.com/run/docs/quickstarts)
[code samples](https://cloud.google.com/run/docs/samples)
[pricing](https://cloud.google.com/run/pricing)
[container runtime contract](https://cloud.google.com/run/docs/container-contract)
[secrets](https://cloud.google.com/run/docs/configuring/secrets)
[vpc](https://cloud.google.com/run/docs/configuring/connecting-vpc#gcloud)
[vpc](https://cloud.google.com/run/docs/securing/private-networking)
[service-to-service auth](https://cloud.google.com/run/docs/authenticating/service-to-service#run-service-to-service-example-go)

````sh
export PROJECT_ID='my-sandbox'
export PROJECT_NUMBER=29220930000
export SERVICE_NAME='my-svc'

gcloud run services list
gcloud run services describe $svc

# deploy
# source: dir or dir with Dockerfile or .tar.gz
#         when provided dir - GCP use Artifact Registry Docker repository
#         and Google Cloud buildpacks to make build.
img="gcr.io/${PROJECT_ID}/${SERVICE_NAME}:latest"
gcloud run deploy $SERVICE_NAME \
  --image=$img \
  --set-env-vars ENV=sandbox \
  --set-env-vars HOST=0.0.0.0 \
  --allow-unauthenticated \
  --platform=managed \
  --tag=traffic-tag-123 \
  --project=$PROJECT_ID

--tag # valid value: [a-z-]

# 0% of traffic
gcloud run deploy --image $img --no-traffic

# 50% of traffic
gcloud run services update-traffic $svc --to-revisions $r=50

# assign 100% of traffic to latest revision
gcloud run services update-traffic $SERVICE_NAME --to-latest

````

Cloud run is per region.
Cloud run configuration has autoscaling with min/max number of instances.
Cloud run URL looks like: `https://$runName-$someHash-uc.a.run.app`.

Cloud Run job - runs container to completion.

To set up custom domain use: global LB, Firebase Hosting, Cloud Run domain mapping.

Startup:
You can enable startup CPU boost to reduce startup latency.
Request waiting for container instance to start
will be kept in a queue for a maximum of 10 seconds.

Processing a request:
For Cloud Run services, CPU is always allocated
when the container instance is processing at least 1 request.

Idle:
CPU depends on "minimum number of container instances" configuration setting.

Shutdown:
When GCP stops app it sends `SIGTERM`.

Cloud Run will only scale out when CPU utilization during request processing exceeds 60%.

Session Affinity - sticky session with TTL for cookie 30 days.
