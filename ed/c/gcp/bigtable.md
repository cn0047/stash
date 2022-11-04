BigTable
-

BigTable - NoSQL database service which scales to patabytes.

4-Dimensional Data Model: row key, column family, column, timestamp.
Data stored as: RowKey:ColumnFamily:Column:Value.

BigTable instance configured per region, and new region can be easily added.

BigTable instance:
* Front-end Server Pool (clients).
* BigTable Cluster (nodes).
* Colossus (tables) - Google's proprietary high durable file system.

BigTable has autoscaling to automatically add and remove nodes.

Replication across clusters works in eventual consistency way.

````sh
gcloud components install cbt

echo project=myprj > ~/.cbtrc

gcloud bigtable clusters list
gcloud bigtable instances list
````
