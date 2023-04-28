RDS (Relational Database Service)
-

````sh
# connect to rds db from local machine:
ssh -i $k -N -L 5431:$rdsHost:$rdsPort ec2-user@$h
ssh -i $k -N -L 5431:$rdsHost:$rdsPort ubuntu@$h
psql -p 5431 -d postgres://$usr:$pwd@localhost/$db
````

Engine:
* Aurora MySQL.
* Aurora PostgreSQL.
* MySQL.
* MariaDB.
* PostgreSQL.
* Oracle.
* Microsoft SQL Server.
