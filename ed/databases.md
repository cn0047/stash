Databases
-

MySQL:

* Transactions.
* Joins (no duplicates).
* Foreign keys, Triggers.
* JSON data type, doesn't support indexing for JSON.
* Proper data indexing can solve the issue with performance.

* *Easy to recover after crash.*

MongoDB:

* Schema-less (perfect for rapid prototyping and for dynamic queries).
* MongoDB don't need ORM.
* Capped collections.
* TTL index.

* *Auto-sharding.*
* *Can be scaled within and across multiple distributed data centers.*
* *Scales easily with no downtime.*

PostgreSQL:

* Transactions.
* Joins.
* Foreign keys, Triggers.
* Supports JSON, XML, HSTORE and Array.

* *Programming Languages support.*

MySQL vs MongoDB:

* MySQL insert bit slower than MongoDB in some cases.
* MySQL selection faster than MongoDB.

MySQL vs PostgreSQL:

* Rich JSON support in PostgreSQL.

PostgreSQL vs MongoDB:

* Selecting, loading, inserting complex document data over 50M records - PostgreSQL faster than MongoDB.
* MongoDB updates JSON data faster than PostgreSQL.
