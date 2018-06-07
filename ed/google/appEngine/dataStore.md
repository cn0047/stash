DataStore ([Google Cloud DataStore](https://cloud.google.com/appengine/docs/standard/go/datastore/))
-

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
