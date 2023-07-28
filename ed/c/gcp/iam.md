IAM (Identity & Access Management)
-

[docs](https://cloud.google.com/iam/docs/)
[permissions](https://cloud.google.com/iam/docs/permissions-reference)
[permissions](https://cloud.google.com/iam/docs/understanding-roles)
[my-groups](https://groups.google.com/my-groups)

````sh
gcloud iam service-accounts list
gcloud iam service-accounts list --format="value(email)"

gcloud iam service-accounts describe fastly@thisisrealtimelog.iam.gserviceaccount.com

gcloud iam workload-identity-pools list --location=global
gcloud iam workload-identity-pools describe $name # name like: projects/123/locations/global/workloadIdentityPools/$pool/providers/$provider

# for GitHub action:
WORKLOAD_IDENTITY_POOL_ID="projects/123/locations/global/workloadIdentityPools/$poolName"
REPO="$gitHubOrg/$gitHubRepoName"
gcloud iam service-accounts add-iam-policy-binding $GCP_SERVICE_ACCOUNT_EMAIL \
  --project="$PROJECT_ID" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/${WORKLOAD_IDENTITY_POOL_ID}/attribute.repository/${REPO}"
# get value for WORKLOAD_IDENTITY_PROVIDER:
gcloud iam workload-identity-pools providers describe $poolProviderName \
  --project="$PROJECT_ID" \
  --location="global" \
  --workload-identity-pool=$poolName \
  --format="value(name)"

````
