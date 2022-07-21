Cloud SQL
-

[doc](https://cloud.google.com/sql/docs/)
[quotas](https://cloud.google.com/sql/docs/quotas)
[quotas mysql](https://cloud.google.com/sql/docs/mysql/quotas)
[pricing](https://cloud.google.com/sql/pricing)
[pricing mysql](https://cloud.google.com/sql/docs/mysql/pricing)
[unsupported statements](https://cloud.google.com/sql/docs/features)
[db flags](https://cloud.google.com/sql/docs/mysql/flags)
[instance](https://cloud.google.com/sql/docs/mysql/instance-settings)
[permissions and roles](https://cloud.google.com/sql/docs/mysql/project-access-control)
[logs](https://console.cloud.google.com/logs/viewer?resource=cloudsql_database)

[connection postgres](https://cloud.google.com/sql/docs/postgres/connect-app-engine)
[connection mysql gen1](https://cloud.google.com/appengine/docs/standard/go/cloud-sql/using-cloud-sql-mysql)
[connection mysql gen2](https://cloud.google.com/appengine/docs/standard/go112/using-cloud-sql)
[connection mysql flex](https://cloud.google.com/appengine/docs/flexible/go/using-cloud-sql)

````bash
# connect to REPL
gcloud sql connect mysql-prod-dbname --user=root
````

````bash
gcloud sql instances list
gcloud sql instances describe $INSTANCE_NAME

gcloud sql databases list --instance=products

gcloud sql users list --instance=products

# connection string
cloudsql(prj:us-central1:dbname)
user:password@cloudsql(prj:us-central1:dbname)

# go111 connection string
unix(/cloudsql/prj:us-central1:instance-name)
user:password@unix(/cloudsql/prj:us-central1:instance-name)/dbname
````

Cloud SQL instance configured per region.
