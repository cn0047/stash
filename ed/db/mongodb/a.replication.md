Replication
-

[FAQ: Replication and Replica Sets](http://docs.mongodb.org/manual/faq/replica-sets/)

Nodes types:
* regular (primary, secondary) [takes part in election]
* arbiter [takes part in election]
* delayed (can't be primary)
* hidden (can't be primary) [takes part in election]

It's possible to use different engines for different nodes!

When node comes back up as a secondary and oplog has looped - the
entire db will be copied from primary.

Minimal Replica Set in mongo:œ
* primary
* secondary
* arbiter (for election when primary down)

````js
rs.initiate({
    "_id": "xmongo",
    "members": [
        {"_id": 0, "host": "xmongo-primary-1:27017", "priority": 10},
        {"_id": 1, "host": "xmongo-secondary-1:27018"},
        {"_id": 2, "host": "xmongo-arbiter-1:27019", "arbiterOnly": true}
    ]
})
// or
rs.initiate()

rs.add('secondary:port')
rs.add('arbiter:port', true) // 2nd parameter - is arbiter

rs.config()
rs.reconfig({})

rs.status()
````

Oplog - (capped collection!!!) commands from master for secondaries.
`--oplogSize 1` = 1MB for oplog.

Chaining - chain of replication: master -> slave 1 -> slave 2 -> slave 3.

#### Durability

Eventual consistency - whether a document was written to all servers
(primaries and secondaries) before control was return to your application.

Durable Replica Set (write concerns):
* Unacknowledged.
* Acknowledged.
* Journaled - wirte data to journal (default).
* Multi-member.

````sh
db.demo.insert({x: 1}, {writeConcern: {w: 2, j: true, wtimeout: 2000}});
# w: 1          // ack from 1 machine
# w: 2          // ack from 2 machines
# w: 'majority' // ack from majority
# j: true       // journaled
````
