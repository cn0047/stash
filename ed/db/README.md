Databases
-

<br>DDL - Data Definition Language: `CREATE, ALTER, DROP, TRUNCATE...`.
<br>DML - Data Manipulation Language: `SELECT, INSERT, UPDATE, DELETE...`.
<br>DCL - Data Control Language: `GRANT, REVOKE`.
<br>TCL - Transaction Control Language: `COMMIT, ROLLBACK, SAVEPOINT...`.

Writes are *40* times more expensive than reads.

Optimize wide. Make writes as parallel as you can.

If table size >100Gb - it should be partitioned.
If table size >1Tb - it should be sharded.

Replication needed for high availability (in case node die - you'll have all data).

When write to two tables (create relation records) wrap it by transaction.

Normalisation not needed when we have:
a very big count of join tables; when data is not updating;

Create table with `uuid()`.

Serializability is an isolation property of transactions, where every transaction may
read and write multiple objects (rows, documents, records).
It guarantees that transactions behave the same as if they had executed
in some serial order (each transaction running to completion before the next transaction starts).

Derived data systems - is the result of taking some existing data from another system
and transforming or processing it in some way.
If you lose derived data, you can recreate it from the original source.
A classic example is a cache.

Transaction - unit of work  performed within DB
that is treated in coherent and reliable way independent of other transactions
(atomic, isolated, with failure recovery & keeps DB consistent).

Referential integrity - constraint on the DB
that makes certain that each foreign key in table point to unique primary key value in another table.

Databases store data either row-oriented or column-oriented.

## ACID

<br>Atomicity - all or nothing (an indivisible and irreducible series of database operations).
<br>Consistency - ensures that any transaction will bring the database from one valid state to another (constraints, cascades, triggers).
<br>Isolation - ensures that the concurrent execution of transactions will executed serially, i.e., one after the other.
<br>Durability - ensures that once a transaction has been committed, it will remain so, even in the event of power loss, crashes, or errors.

## BASE

BASE is opposite to ACID.

<br>Basically
<br>Available - basic reading and writing operations available as much as possible, but without any kind of consistency guarantees.
<br>Soft state - without consistency guarantees, after some amount of time, we only have some probability of knowing the state.
<br>Eventually consistent - after some time we will eventually be able to know what the state of the database is.

## Two-Phase Locking (2PL)

2 phases:
* Expanding phase: locks are acquired and no locks are released.
* Shrinking phase: locks are released and no locks are acquired.

Locks types:
* Read-lock, Shared.
* Write-lock, Exclusive.

<br>Acquired Read-lock  blocks new Write-lock.
<br>Acquired Write-lock blocks new Read-lock.
<br>Acquired Write-lock blocks new Write-lock.

Mutual blocking between transactions results in a deadlock.

## Two-Phase Commit (2PC)

Update account table with transaction.

0. [account table]      Add field pendingTransactions.
1. [transactions table] Insert record with state initial.
2. [transactions table] Update record with state initial to pending.
3. [account table]      Update record (decrement) & add transaction to pendingTransactions field for this record.
4. [account table]      Update next record (increment) & add transaction to pendingTransactions field for this record.
5. [transactions table] Update transaction with state pending to committed.
6. [account table]      Delete transactions from pendingTransactions.
7. [transactions table] Update transaction set state done.

## MVCC (Multiversion concurrency control)

It's a concurrency control method in DBs
which provide concurrent access to DB and transactions in programming languages.

MVCC keep multiples copies of each data item,
in this way, each user connected to DB sees a snapshot of the DB at a particular instant in time.
<br>Any changes made by a writer will not be seen by other users of the DB
until the changes have been completed (committed transaction).
<br>When an MVCC DB needs to update an item of data, it will not overwrite the original data item with new data,
but instead creates a newer version of the data item. Thus there are multiple versions stored.
<br>The version that each transaction sees depends on the isolation level implemented.

## Sharding

Problems With Sharding:
* Imagine that shard outgrows storage capacity and must be splited (rebalancing data).
* Join data from multiple shards - may turn out to be a problem.

## Replication

* Statement Based Replication.
* Row Based Replication.

* Based on Write-Ahead Log - changes are first recorded in the log,
  and nodes (master and slave) processes this log and builds exact same data.
* Trigger Based Replication - triggers write changes to replication table
  and separate process handle that changes. Big overhead and prone to bugs.

Read-after-write Consistency:
Always read the user’s own profile from the master and any other users’ profiles from a slave,
because due to replication lag user may not see recent updates on slave.

## Distributed transaction

* Don't use distributed transactions, it won't scale - have to use eventual consistency.
* Rely on idempotency.
* SAGA.

## Saga Pattern

Saga is a sequence of transactions where each transaction
updates data within a single service,
1st transaction is initiated by an external request
and each subsequent step is triggered by the completion
of the previous one.

Implementations:
* Choreography (simplicity, loose coupling; difficult to understand, cyclic dependencies between services).
* Orchestration (all complexity in 1 place).

Anomalies:
* Lost updates — one saga overwrites without reading changes made by another saga.
* Dirty reads — one saga reads updates that another saga hasn't yet completed.
* Fuzzy/nonrepeatable reads — two sagas read same data and get different results because another saga has made updates.

Countermeasures:
* Semantic lock - sets flag in record which indicates that record isn’t committed yet.
* Commutative updates - operations are commutative if they can be executed in any order (debit, credit).
* Pessimistic view - reorders steps of saga to minimize business risk due to dirty read.
* Reread value - rereads record before updating it, verifies it’s unchanged, then updates record.
* Version file - handle out-of-order requests (ensure AuthorizeCard before CancelAuthorization).
* By value - selecting concurrency mechanisms based on business risk.

Issues:
* During rollback when some subsequent rollback fails.

## Database selection

Questions to answer:
* How is the data structured (unstructured, key-value, semi-structured, relational)?
* Is ACID compliance required?
* What consistency model is required (strong, eventual)?
* What data types and sizes required (text, spatial, time-series, BLOB)?
* How will scalability requirements change over time?
* What is read/write queries ratio?
* How is data sensitive (protection, encryption)?
* What level of durability is required?
