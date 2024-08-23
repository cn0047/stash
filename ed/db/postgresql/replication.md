Replication
-

Physical replication - uses exact block addresses and byte-by-byte replication.

Logical replication - method of replicating data objects and their changes,
based upon their replication identity (usually a primary key).

Master-Master replication can be implemented via logical replication (with conflict management).
Logical replication doesn't handle conflicts automatically, it must be implemented on application-level
or via trigger-based conflict resolution strategies.

## Streaming replication

Primary (sending) server sends data.
Standby server always receives replicated data.

## Logical replication

Publisher server (`CREATE PUBLICATION`) replicate data to subscriber.
Subscriber server (`CREATE SUBSCRIPTION`) always receives replicated data.
