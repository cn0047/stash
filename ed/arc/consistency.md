Consistency models
-

**Strong (Immediate) consistency**:
Strongest consistency model.
Accesses to the item after an update will be consistent for all processes.

**Sequential consistency**:
Item write doesn't have to be seen instantaneously,
but, writes to item by different processes have to be seen in the same order by all processes.

**Causal consistency**:
Weakening model of sequential consistency.
Only write operations that are causally related, need to be seen in the same order by all processes.
(Client is guaranteed to read at least its own writes.)

**Eventual consistency (Optimistic Replication)**:
For high availability.
Guarantees that, eventually all accesses to the item will return the last updated value.

**Strong eventual consistency**:
Guarantees that any two nodes that have received the same set of updates will be in the same state.

## Linearizability (Atomic consistency)

DB could give the illusion that there is only one replica - this is the idea behind linearizability,
AKA atomic|strong|immediate|external consistency.

In a linearizable system, as soon as one client successfully completes a write,
all clients reading from the database must be able to see the value just written.

The basic idea behind linearizability is simple:
to make a system appear as if there is only a single copy of the data.

Implementing Linearizable Systems:
* Single-leader replication (potentially linearizable).
* Consensus algorithms (linearizable).
* Multi-leader replication (not linearizable).
* Leaderless replication (probably not linearizable).

For linearizability important ordering guarantees so it may be confused with serializability,
but they are different,
serializability related to db transaction isolation level,
but linearizability is a recency guarantee on reads and writes of a register (an individual object),
moreover linearizability guarantees order, like: `a -> b -> c`
but for serializability it's ok when: `b -> a -> c` only important separately execute actions.
