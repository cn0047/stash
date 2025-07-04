History of PostgreSQL
-

[info](https://www.postgresql.org/docs/release/)

#### 17

* New memory management for VACUUM reduces memory usage and boosts vacuum performance.
* Enhanced SQL/JSON support: constructors, identity funcs, and JSON_TABLE().
* Added incremental backup capability with pg_basebackup.

#### 16

* Added parallel support for FULL and right‑outer hash joins.
* Logical replication now possible from standby servers.
* Logical subscribers can apply large transactions in parallel.
* SQL/JSON constructors and identity functions added.
* Improved vacuum freezing performance.
* New pg_stat_io view enables I/O stats monitoring.

#### 15

* Implemented SQL-standard MERGE statement.
* Improved sort algorithms and performance.
* Server logs can output in JSON format.
* Public schema now owned by database owner with tighter permissions.

#### 14

* Extended statistics on expressions for better query plans.
* Introduced predefined roles: pg_database_owner, pg_read_all_data, pg_write_all_data.
* Long-running queries auto-cancel if client disconnects.

#### 13

* Imporoved B-tree deduplication.
* Imporoved performance for aggregates and partitioned queries.
* Added support for parallel vacuuming of indexes and incremental sort.

#### 12

* Added support for SQL/JSON path expression.
* Added support for generated (computed) columns.
* Introduced pluggable table storage interface.
* Improved query and storage performance.

#### 11

* Added support for stored procedures with transaction control.
* Improved partitioning performance.

#### 10

* Declarative table partitioning.
* Added support for logical replication.
* Improved query parallelism.

#### 9.0

* Introduced built-in streaming replication for high availability.
* Introduced hot standby for read-only queries on replicas.
* Introduced in-place upgrade (pg_upgrade).

#### 8.0

* Savepoints support for nested transactions.
* Inroduced tablespaces and point-in-time recovery.

#### 7.0

* Added support for foreign keys.

#### 6.0

* Added support for unique index.
* Introduced pg_dumpall utility.
