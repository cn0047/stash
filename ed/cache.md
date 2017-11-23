cache
-

Memcached:

* In-memory key-value store.
* Can store string up to 1MB.

Redis:

* In-memory data structure store, used as database.
* Can store string up to 512MB.
* Transactions - yes.
* Durability - yes.
* Server-side scripts - lua.

* *Partitioning methods - sharding.*
* *Replication methods - master-slave replication.*

Redis vs Memcached:

1. In operations with sets Redis exceed Memcached.
2. In operations with gets Redis loses to Memcached.
3. Memcached faster than Redis with multi gets at high volumes (more than 100K entries).
4. Redis and Memcached are similar on multi gets for small volumes (1000 or 100 entries).
