Engine
-

````
mongod -storageEngine wiredTiger
````

#### In-Memory Storage Engine (MongoDB Enterprise)

#### Wired Tiger (default since v3.2)

* Document level lock.
* Compression (data, indexes).
* No in place updates.

#### MMAPv1

* Collection level lock.
* In place update
  (move doc in memory to place where present sufficient space for update).
* Uses `Power of 2 Sized Allocations`
  (document stored as record with document itself and extra space, or padding).
