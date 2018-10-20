DataStore
-

[doc](https://cloud.google.com/appengine/docs/standard/go/datastore/)
[backups](https://cloud.google.com/appengine/articles/scheduled_backups)
[console](https://console.cloud.google.com/datastore/)
[limits](https://cloud.google.com/datastore/docs/concepts/limits)

````bash
gcloud datastore cleanup-indexes

gcloud datastore operations list
````

````bash
curl -x GET  'https://datastore.googleapis.com/$discovery/rest?version=v1' | jq
curl -X POST 'https://datastore.googleapis.com/v1/projects/itismonitoring:export'
````

## Overview

Google Cloud DataStore is a NoSQL document database.
In DataStore nested transactions are not supported.

Cloud Datastore us
es optimistic concurrency to manage transactions.

````
Relational DB ⇒ Table ⇒ Row    ⇒ Field    ⇒ Primary key
Datastore     ⇒ Kind  ⇒ Entity ⇒ Property ⇒ Key
````

When you create an entity, you can optionally designate another entity as its parent.
An entity without a parent is a root entity.

A transaction on entities belonging to different entity groups is called a cross-group (XG) transaction.
The transaction can be applied across a maximum of twenty-five entity groups.
In a single-group transaction, you cannot perform a non-ancestor query in an XG transaction.

Comparison operators: `=, <, <=, >, >=`.

An entity is limited to 1 megabyte when stored.
That roughly corresponds to a limit of 1 megabyte for the serialized form of this message.

`Megastore` - the underlying technologies for datastore.

Datastore contention occurs when a single entity or entity group is updated too rapidly.
The datastore will queue concurrent requests to wait their turn.
Requests waiting in the queue past the timeout period will throw a concurrency exception.

#### Transaction

Isolation levels (from strongest to weakest):

* Serializable
* Repeatable Read
* Read Committed
* Read Uncommitted

When a commit returns successfully, the transaction is guaranteed to be applied,
but that does not mean the result of your write is immediately visible to readers.
Applying a transaction consists of two milestones:

* Milestone A – the point at which changes to an entity have been applied
* Milestone B – the point at which changes to indices for that entity have been applied

There are two major reasons that datastore transaction errors occur:
* Timeouts due to write contention
  (when you attempt to write to a single entity group too quickly)
* Timeouts due to datastore issues
  (due to the distributed nature of Bigtable, which the datastore is built on)

Whenever possible, make your datastore transactions idempotent
so that if you repeat a transaction, the end result will be the same.

#### Index

A query can't find property values that aren't indexed, nor can it sort on such properties.

Only indexed properties can be projected. The same property cannot be projected more than once.

`Built-in index` - by default,
datastore automatically predefines an index for each property of each entity kind.

`Composite index` - index stored in `index.yaml`.

The total number of indexes is 2^(number of filters) * (number of different orders),
for example: 2 ^ 5 * 4 = 128 indexes

In case index does not exist - the datastore returns a NeedIndexError.

````sh
rm index.yaml
gcloud vacuum_indexes
gcloud update_indexes
````

#### GQL

````sql
select * from Tweet where __key__ = Key(Tweet, "id-123")
````

#### +/-

Advantages:
* Transactions.

Disadvantages:
* Do not support substring matches, case-insensitive matches, full-text search.
* Do not support `NOT, !=, OR, IN`.
* No REPL for dev environment.
* No delete query in prod REPL, [look](https://monosnap.com/file/0osxGC8ocSQrQxGPx05ByYz3FkNByh).
* No count query in prod REPL.
* Only one inequality filter per query is supported. Encountered both ScheduledDate and Updated.
* No way to drop kind (collection).
* No way to re-use compound index (add 1 extra field to existing index - gae will generate new index).
* If `index.yaml` not uploaded and field don't have index - no possibility to run query against this index.
* `select template from Article` - won't work because no index.
* Property body is too long. Maximum length is 1500. (`noindex` can fix this issue).
* [How delete element from array](https://monosnap.com/file/YrQHARwcRPAEagaNfoKeMhh1o1bsnZ).
* How reproduce prod datastore timeout problem?
* DataStore: cannot load field "id" into a "Project.Entity": no such struct field
  (after rename field in struct when have entities in datastore with property id).

#### GO

````go
key := datastore.NewKey(
        ctx,        // context.Context
        "Employee", // Kind
        "asalieri", // String ID; empty means no string ID
        0,          // Integer ID; if 0, generate automatically. Ignored if string ID specified.
        nil,        // Parent Key; nil means no parent
)

_, err = datastore.PutMulti(ctx, []*datastore.Key{k1, k2, k3}, []interface{}{e1, e2, e3})

var entities = make([]*T, 3)
err = datastore.GetMulti(ctx, []*datastore.Key{k1, k2, k3}, entities)

err = datastore.DeleteMulti(ctx, []*datastore.Key{k1, k2, k3})

q := datastore.NewQuery("Person").KeysOnly()

q := datastore.NewQuery("User").Filter("Tags =", "test").Filter("Tags =", "go").Order("-Name")
u := make([]User, 0)
_, err := q.GetAll(ctx, &u)

// not equal
WHERE tags >= 'math' AND tags <= 'math'
````
