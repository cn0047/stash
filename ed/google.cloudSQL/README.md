Cloud SQL
-

[quotas](https://cloud.google.com/sql/docs/quotas)
[pricing](https://cloud.google.com/sql/pricing)
[unsupported statements](https://cloud.google.com/sql/docs/features)
[db flags](https://cloud.google.com/sql/docs/mysql/flags)

````bash
# in Cloud Shell
gcloud sql connect mysql-prod-hera --user=root
````

````bash
gcloud sql instances list

gcloud sql databases list --instance=mysql-prod-hera
````
