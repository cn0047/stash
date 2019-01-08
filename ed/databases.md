Databases
-

## MVCC (Multiversion concurrency control):

It's a concurrency control method in DBs
which provide concurrent access to DB and transactions in programming languages.

MVCC keep multiples copies of each data item,
in this way, each user connected to DB sees a snapshot of the DB at a particular instant in time.
<br>Any changes made by a writer will not be seen by other users of the DB
until the changes have been completed (committed transaction).
<br>When an MVCC DB needs to update an item of data, it will not overwrite the original data item with new data,
but instead creates a newer version of the data item. Thus there are multiple versions stored.
<br>The version that each transaction sees depends on the isolation level implemented.

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

* *Programming Languages support.*

## MySQL vs MongoDB:

* MySQL insert bit slower than MongoDB in some cases.

* *MySQL selection faster than MongoDB.*

## MySQL vs PostgreSQL:

* Rich JSON support in PostgreSQL (MySQL doesn't support indexing for JSON).

## PostgreSQL vs MongoDB:

* Selecting, loading, inserting complex document data over 50M records - PostgreSQL faster than MongoDB.
* MongoDB updates JSON data faster than PostgreSQL.
