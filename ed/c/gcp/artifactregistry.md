Artifact Registry
-

[docs](https://docs.cloud.google.com/artifact-registry/docs)

````sh
a="Authorization: Bearer $(gcloud auth print-access-token)"
jh="Content-Type: application/json"

prj=''
curl -X POST -H $a -H $jh  "https://artifactregistry.googleapis.com/v1/projects/$prj/locations/us/repositories/npm:testIamPermissions" -d '
  {"permissions":["artifactregistry.repositories.uploadArtifacts","artifactregistry.repositories.downloadArtifacts","artifactregistry.repositories.get"]}
'
````
