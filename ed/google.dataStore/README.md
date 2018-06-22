DataStore
-

[doc](https://cloud.google.com/appengine/docs/standard/go/datastore/)

Google Cloud DataStore is a NoSQL document database.
In DataStore nested transactions are not supported.

Cloud Datastore uses optimistic concurrency to manage transactions.

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

#### +/-

Advantages:
* Transactions.

Disadvantages:
* `FIND ALL WHERE id IN (1, 2)`, `FIND ALL WHERE id = 1 OR id = 2`.
* [How delete element from array](https://monosnap.com/file/YrQHARwcRPAEagaNfoKeMhh1o1bsnZ).
* Only one inequality filter per query is supported. Encountered both ScheduledDate and Updated.

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
````
