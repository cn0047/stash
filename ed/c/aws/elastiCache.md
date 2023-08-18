Elasti Cache
-

Engines: redis, memcached.

Redis/Memcached cluster per region and per VPC.

````sh
aws --profile=$p elasticache describe-cache-clusters
aws --profile=$p elasticache describe-replication-groups
aws --profile=$p elasticache describe-snapshots
````
