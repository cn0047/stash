Comparison
-

## MySQL:

* Transactions.
* Joins (no duplicates - less disc space).
* Foreign keys, Triggers.
* JSON data type.

* *Easy to recover after crash.*

## MongoDB:

* Schema-less (perfect for rapid prototyping and for dynamic queries).
* TTL index.
* Auto-sharding.
* ↓ MongoDB don't need ODM (ORM).
* ↓ Capped collections.

* *Can be scaled within and across multiple distributed data centers.*
* *Scales easily with no downtime.*

## PostgreSQL:

* Transactions.
* Joins.
* Foreign keys, Triggers.
* Supports JSON, XML, **HSTORE** and Array.
* ↓ Inheritance.
* ↓ Parallel query.

* *Programming Languages support.*

## Comparison:

MySQL vs MongoDB:
* MySQL insert bit slower than MongoDB in some cases.
* *MySQL selection faster than MongoDB.*

MySQL vs PostgreSQL:
* Rich JSON support in PostgreSQL (MySQL doesn't support indexing for JSON).

PostgreSQL vs MongoDB:
* Selecting, loading, inserting complex document data over 50M records - PostgreSQL faster than MongoDB.
* MongoDB updates JSON data faster than PostgreSQL.
