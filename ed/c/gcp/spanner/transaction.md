Transaction
-

[docs](https://cloud.google.com/spanner/docs/transactions)

Spanner provides clients with the strictest concurrency-control guarantees for transactions - external consistency.
TrueTime is a highly available, distributed clock that is provided to applications on all GCP servers.

Avoid mixing DML and mutation in the same transaction.

Transaction modes:
* Locking **read-write** - writing data into Spanner.
  Rely on pessimistic locking and, if necessary, two-phase commit.
* **Read-only** - provides guaranteed consistency across several reads.
  It might wait for in-progress writes to complete before executing.
* Partitioned DML - designed for bulk updates and deletes.

A read-write transaction in Spanner executes a set of reads and writes atomically
at a single logical point in time (ACID).

Read types:
* strong (default) - see all data that has been committed up until the start of this read.
* stale - read at a timestamp in the past.

Locks are taken at the granularity of row-and-column.

When deadlock conditions met - spanner forces 1 client to abort.

Reads:

[docs](https://cloud.google.com/spanner/docs/reads#go)

````sh
# read
stmt := spanner.Statement{SQL: `SELECT SingerId, AlbumId, AlbumTitle FROM Albums`}
iter := client.Single().Query(ctx, stmt)

# using index
iter := client.Single().ReadUsingIndex(
  ctx, "Albums", "AlbumsByAlbumTitle", spanner.AllKeys(), []string{"AlbumId", "AlbumTitle"}
)

# strong read
iter := client.Single().Read(ctx, "Albums", spanner.AllKeys(), []string{"SingerId", "AlbumId", "AlbumTitle"})

# stale read
ro := client.ReadOnlyTransaction().WithTimestampBound(spanner.ExactStaleness(15 * time.Second))
defer ro.Close()
iter := ro.Read(ctx, "Albums", spanner.AllKeys(), []string{"SingerId", "AlbumId", "AlbumTitle"})
defer iter.Stop()

````
