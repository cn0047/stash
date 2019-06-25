Databases
-

Writes are 40 times more expensive than reads.

Optimize wide. Make writes as parallel as you can.

Replication needed for high availability (in case node die - you'll have all data).

## Sharding

Problems With Sharding:
* Imagine that shard outgrows storage capacity and must be splited (rebalancing data).
* Join data from multiple shards - may turn out to be a problem.

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

## Two-Phase Locking (2PL)

2 phases:
* Expanding phase: locks are acquired and no locks are released.
* Shrinking phase: locks are released and no locks are acquired.

Locks types:
* Read-lock, Shared
* Write-lock, Exclusive

<br>Acquired Read-lock  blocks new Write-lock.
<br>Acquired Write-lock blocks new Read-lock.
<br>Acquired Write-lock blocks new Write-lock.

Mutual blocking between transactions results in a deadlock.

## Two-Phase Commit (2PC)

Update account table with transaction.

0. Add field pendingTransactions into account table.
1. In table transactions insert record with state initial.
2. Update record with state initial to pending in transactions table.
3. Update record in account table (decrement)
   & add transaction to pendingTransactions field for this record.
4. Update next record in account table (increment)
   & add transaction to pendingTransactions field for this record.
5. Update transaction in transactions table with state pending to committed.
6. Update account table delete transactions from pendingTransactions.
7. Update transaction in transactions table set state done.

## Saga Pattern

Saga is a sequence of local transactions where each transaction
updates data within a single service,
1st transaction is initiated by an external request
and each subsequent step is triggered by the completion
of the previous one.

implementations: Choreography, Orchestration.

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

## MySQL vs MongoDB:

* MySQL insert bit slower than MongoDB in some cases.

* *MySQL selection faster than MongoDB.*

## MySQL vs PostgreSQL:

* Rich JSON support in PostgreSQL (MySQL doesn't support indexing for JSON).

## PostgreSQL vs MongoDB:

* Selecting, loading, inserting complex document data over 50M records - PostgreSQL faster than MongoDB.
* MongoDB updates JSON data faster than PostgreSQL.
