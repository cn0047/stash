Partitioning
-

[docs](https://www.postgresql.org/docs/current/ddl-partitioning.html)

Forms of partitioning:
* Range Partitioning - range from 1 to 10, from 10 to 20, etc.
* List Partitioning - explicit keys list.
* Hash Partitioning - hash value of the partition key divided by the specified modulus.

The partitioned table itself - "virtual" table having no storage of its own.
It is not possible to turn regular table into partitioned table or vice versa.
Partitions can also be foreign tables, although considerable care is needed.

````sql
ATTACH PARTITION
DETACH PARTITION

````
