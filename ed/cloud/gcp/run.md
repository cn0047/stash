Cloud Run
-

[pricing](https://cloud.google.com/run/pricing)

````sh
export PROJECT_ID='my-sandbox'
export PROJECT_NUMBER=29220930000
export SERVICE_NAME='my-svc'

# deploy
# source: dir or dir with Dockerfile or .tar.gz
#         when provided dir - GCP use Artifact Registry Docker repository
#         and Google Cloud buildpacks to make build.
img="gcr.io/${PROJECT_ID}/${SERVICE_NAME}:latest"
gcloud run deploy $SERVICE_NAME \
  --image=$img \
  --set-env-vars ENV=sandbox \
  --set-env-vars HOST=0.0.0.0 \
  --allow-unauthenticated --platform=managed \
  --project=$PROJECT_ID

# assign 100% of traffic to latest revision
gcloud run services update-traffic $SERVICE_NAME --to-latest

````
