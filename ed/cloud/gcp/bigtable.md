BigTable
-

BigTable - is NoSQL and it scales to patabytes.

4-Dimensional Data Model: row key, column family, column, timestamp.
Data stored as: RowKey:ColumnFamily:Column:Value.

BigTable instance:
* Front-end Server Pool (clients).
* BigTable Cluster (nodes).
* Colossus (tables) - Google's proprietary high durable file system.

Replication across clusters works in eventual consistency way.

````sh
gcloud components install cbt

echo project=myprj > ~/.cbtrc
````
