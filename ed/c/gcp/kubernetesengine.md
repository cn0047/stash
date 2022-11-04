GKE - Google Kubernetes Engine
-

[docs](https://cloud.google.com/kubernetes-engine/docs/)
[pricing](https://cloud.google.com/kubernetes-engine/pricing)

````sh
gcloud container clusters list # --project $p
gcloud container clusters get-credentials
gcloud container clusters get-credentials $name # --project $p --region $r

````

GKE may be per zone or per region.
