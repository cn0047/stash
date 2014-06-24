Replication
-

*MySQL 5.5*

Depending on the configuration, you can replicate all databases, selected databases, or even selected tables within a database.
Asynchronous replication - one server acts as the master, while one or more other servers act as slaves.
Synchronous replication which is a characteristic of MySQL Cluster.
There are two core types of replication format:
* Statement Based Replication (SBR) - which replicates entire SQL statements.
* Row Based Replication (RBR) - which replicates only the changed rows.
