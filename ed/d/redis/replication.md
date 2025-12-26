Replication
-

Replication can be used for scalability, to have multiple replicas for read-only queries operations,
or for improving data safety and high availability.

Redis replication based on leader follower (master-replica) replication.
Redis uses by default asynchronous replication, for low latency and high performance.
Synchronous replication of certain data can be requested by the clients using the WAIT command.
Redis replication is non-blocking on master side.
Replication is also largely non-blocking on the replica side.

In setups where redis replication is used, it is strongly advised to have persistence turned on
in the master and in the replicas.
