Engine
-

````
mongod -storageEngine wiredTiger
````

#### MMAPv1

* Collection level lock.
* In place update
  (move doc in memory to place where present sufficient space for update).
* Power of two sizes.

#### Wired Tiger

* Document level lock.
* Compression (data, indexes).
* No in place updates.
