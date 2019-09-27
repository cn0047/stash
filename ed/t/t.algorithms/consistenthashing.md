Consistent Hashing
-

Consistent Hashing is a distributed hashing scheme
that operates independently of the number of servers or objects in a distributed hash table
by assigning them a position on an abstract circle (HashRing).

1. Create HashRing of all possible values.
   (You can think of the circle as all integers 0 ..2^32-1.)
2. Place servers on the HashRing.
3. Read/Write data:
  1. Calculate the hash `serverIndex = hashValue % n` and map it to some position on the HashRing.
  2. If server there - ok, otherwise travel clockwise on the ring until find the first server.

When a hash table is resized, only `keys/n` need to be remapped.
