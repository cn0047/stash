Array
-

Find the number which occurs odd number times in an Array:
`XOR` all values in array.

Rotated array - array where 2 parts sorted but this parts not consequent.
Pivot element of array - element which changes the order of array.

Sub array - set of contiguous elements.

Majority element - number of occurrences > n/2.

## Array types

**Bit array** (bitmap, bitset, bit string, or bit vector) - is array data structure that compactly **stores bits**.
Is effective at exploiting bit-level parallelism in hardware to perform operations quickly.

**Circular buffer** (circular queue, cyclic buffer or ring buffer) -  is data structure
that uses a single, fixed-size buffer as if it were **end-to-end connected**.

**Dynamic array** (growable array, resizable array, dynamic table, mutable array, or array list) -
is a **random access, variable-size** list data structure that allows elements to be added or removed.

**Sparse array** - is **an array** in which most of the elements have the **default value** (usually 0 or null).

**Hashed array tree** - is dynamic array data-structure which have
top-level array where each entry is pointer to array with elements.

**Hash table** (hash map) - is a data structure
used to implement an associative array, a structure that can **map keys to values**.
A hash table uses a hash function
to compute an index into an array of buckets or slots, from which the desired value can be found.
The average cost for each lookup is independent of the number of elements stored in the table.
Many hash table designs also allow arbitrary insertions and deletions of key-value pairs.
Widely used in many kinds of computer software,
particularly for associative arrays, database indexing, caches, and sets.
Examples:
  * ASCII implementation = sum ASCII's of string. Order in string don't effect hash (foo & oof have same hash).
  * CRC32 - not secure.
  * MD5 - not efficient (slow), not secure.
  * SHA-2 - not efficient.
Dealing with collisions:
  1. Open addressing - if hash is occupied -> hash++ and try again.
  2. Chaining (simpler to support) - if hash is occupied -> create linked list from old hash value and new hash value.
Collisions depends on how many free slots available and how much slots populated (fill factor).
