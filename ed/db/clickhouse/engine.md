Engine
-

#### MergeTree

MergeTree - supports data replication, partitioning, secondary data-skipping indexes, and more.

Engines in the family:
* MergeTree - for high data ingest rates and huge data volumes.
* ReplacingMergeTree - differs from MergeTree in that it removes duplicate entries with same sorting key.
* SummingMergeTree - replaces rows with same primary key (or sorting key) with one row which contains summarized values.
* AggregatingMergeTree - replaces rows with same primary key (or sorting key) with one row which contains combination of states of aggregate functions.
* CollapsingMergeTree - inherits from MergeTree and adds logic for collapsing rows during the merge process.
* VersionedCollapsingMergeTree - same as CollapsingMergeTree but uses different collapsing algorithm.
* GraphiteMergeTree.

#### Log

Log - lightweight engines with minimum functionality.

Engines in the family:
* TinyLog.
* StripeLog.
* Log.

#### Integration engine

Integration Engines - Engines for communicating with other data storage and processing systems.

Engines in the family:
* ODBC.
* JDBC.
* MySQL.
* MongoDB.
* Redis.
* HDFS.
* S3.
* Kafka.
* EmbeddedRocksDB.
* RabbitMQ.
* PostgreSQL.
* S3Queue.

#### Special

Engines in the family:
* Distributed.
* Dictionary.
* Merge.
* File.
* Null.
* Set.
* Join.
* URL.
* View.
* Memory.
* Buffer.
* KeeperMap.

#### Virtual columns

Virtual column - integral table engine attribute that is defined in engine source code.
