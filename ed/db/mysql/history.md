History of MySQL
-

#### MySQL 8

* Invisible index.
* Algorithm=instant.

#### MySQL 5.7

* **JSON functions.**
* Security improvements.
* SQL mode changes.
* Online ALTER TABLE (RENAME INDEX).
* InnoDB enhancements.
* Condition handling.
* Optimizer.
* **Triggers (Multiple triggers are permitted).**
* Logging.
* Test suite.
* mysql client.
* Database name rewriting with mysqlbinlog.
* HANDLER with partitioned tables.
* Index condition pushdown support for partitioned tables.
* WITHOUT VALIDATION support for ALTER TABLE ... EXCHANGE PARTITION.
* Master dump thread improvements.
* Globalization improvements.
* **Changing the replication master without STOP SLAVE.**
* **Deprecated EXPLAIN EXTENDED and PARTITIONS.**
* INSERT DELAYED is no longer supported.

#### MySQL 5.6

* Security improvements.
* MySQL Enterprise.
* Changes to server defaults.
* **InnoDB enhancements** (FULLTEXT indexes, ALTER TABLE without blocking, DATA DIRECTORY clause of the CREATE TABLE (which allows you to create InnoDB file-per-table tablespaces)).
* Partitioning (Maximum number of partitions is increased to 8192).
* Performance Schema.
* MySQL Cluster.
* Replication and logging (Transaction-based replication (Synchronous replication)).
* **Optimizer enhancements** (EXPLAIN for DELETE, INSERT, REPLACE, UPDATE.).
* Condition handling.
* Data types (Permits fractional seconds for TIME, DATETIME, and TIMESTAMP).
* Host cache.
* OpenGIS.
* **Deleting index don't block read and write**.
* **Index usage statistics**.

#### MySQL 5.5

* MySQL Enterprise Thread Pool.
* MySQL Enterprise Audit (audit_log).
* Pluggable authentication.
* Multi-core scalability.
* InnoDB I/O subsystem.
* Diagnostic improvements.
* Solaris.
* **Default storage engine (Is InnoDB).**
* MySQL Cluster.
* **Semisynchronous replication.**
* Unicode.
* Partitioning (RANGE COLUMNS, LIST COLUMNS, and else).
* SIGNAL and RESIGNAL.
* Metadata locking.
* IPv6 support.
* XML (LOAD XML INFILE).
* Build configuration.
* **Added `utf8mb4` character set** (utf8 - 3 bytes per character, utf8mb4 - 4).

#### MySQL 5.1

* **Partitioning.**
* **Row-based replication.**
* Plugin API.
* **Event scheduler.**
* Server log tables.
* Upgrade program.
* MySQL Cluster.
* Backup of tablespaces.
* Improvements to INFORMATION_SCHEMA.
* XML functions with XPath support.
* Load emulator ([mysqlslap](http://dev.mysql.com/doc/refman/5.1/en/mysqlslap.html) program).

#### MySQL 5.0

* Information Schema.
* Instance Manager.
* Precision Math.
* Storage Engines.
* **Stored Routines.**
* **Triggers.**
* Views.
* Cursors.
* Strict Mode and Standard Error Handling.
* VARCHAR Data Type.
* BIT Data Type.
* Optimizer enhancements.
* XA Transactions.
