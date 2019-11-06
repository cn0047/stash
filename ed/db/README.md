Databases
-

Writes are *40* times more expensive than reads.

Optimize wide. Make writes as parallel as you can.

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

<br>DDL - Data Definition Language: `CREATE, ALTER, DROP, TRUNCATE...`.
<br>DML - Data Manipulation Language: `SELECT, INSERT, UPDATE, DELETE...`.
<br>DCL - Data Control Language: `GRANT, REVOKE`.
<br>TCL - Transaction Control Language: `COMMIT, ROLLBACK, SAVEPOINT...`.

## ACID

<br>Atomicity - all or nothing.
<br>Consistency - ensures that any transaction will bring the database from one valid state to another (constraints, cascades, triggers).
<br>Isolation - ensures that the concurrent execution of transactions will executed serially, i.e., one after the other.
<br>Durability - ensures that once a transaction has been committed, it will remain so, even in the event of power loss, crashes, or errors...

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

Saga is a sequence of local transactions where each transaction
updates data within a single service,
1st transaction is initiated by an external request
and each subsequent step is triggered by the completion
of the previous one.

implementations: Choreography, Orchestration.

## Linearizability

DB could give the illusion that there is only one replica - this is the idea behind linearizability,
AKA atomic|strong|immediate|external consistency.

In a linearizable system, as soon as one client successfully completes a write,
all clients reading from the database must be able to see the value just written.

The basic idea behind linearizability is simple:
to make a system appear as if there is only a single copy of the data.
