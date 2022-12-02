liquibase
-

[docs](https://docs.liquibase.com/)

````sh
docker run -it --rm -v $PWD:/app -w /app liquibase/liquibase /bin/bash

liquibase --version

````

Examples:

````sh
# GCP Spanner
cd ed/sh/liquibase
u="jdbc:cloudspanner://$SPANNER_EMULATOR_HOST/projects/$prj/instances/$dbi/databases/$db;usePlainText=true"
cl=examples/gcpSpanner/changelog.xml
cp=examples/gcpSpanner/liquibase-spanner-4.5.0-all.jar

liquibase --url=$u --changelog-file=$cl --classpath=$cp --hub-mode=off updateSQL
liquibase --url=$u --changelog-file=$cl --classpath=$cp --hub-mode=off update
liquibase --url=$u --changelog-file=$cl --classpath=$cp --hub-mode=off rollbackCountSQL 1
liquibase --url=$u --changelog-file=$cl --classpath=$cp --hub-mode=off rollbackCount 1

````
