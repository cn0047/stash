Cloud SQL
-

[quotas](https://cloud.google.com/sql/docs/quotas)
[pricing](https://cloud.google.com/sql/pricing)
[unsupported statements](https://cloud.google.com/sql/docs/features)
[db flags](https://cloud.google.com/sql/docs/mysql/flags)
[instance](https://cloud.google.com/sql/docs/mysql/instance-settings)
[permissions and roles](https://cloud.google.com/sql/docs/mysql/project-access-control)

````bash
# in Cloud Shell
gcloud sql connect mysql-prod-hera --user=root
````

````bash
gcloud sql instances list
gcloud sql instances describe $INSTANCE_NAME

gcloud sql databases list --instance=mysql-prod-hera
````
