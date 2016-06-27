Data structures
-

#### Abstract

* Container - is a class, a data structure, or an abstract data type (ADT)
whose instances are collections of other objects. In other words, container store **objects in an organized way**.

* Collection - is a grouping of some variable number of data items (possibly zero)
that have some shared significance to the problem being solved (Generally, the data items will be of the **same type**).

* Associative array, map, symbol table, or dictionary - is an abstract data type
composed of a collection of **(key, value) pairs**, such that each possible **key appears just once** in the collection.

* Multimap generalizes an associative array by allowing **multiple values to be associated with a single key**.

* Bidirectional map - is a related abstract data type in which the **bindings operate in both directions**:
each value must be associated with a unique key,
and a second lookup operation takes a value as argument and looks up the key associated with that value.

* List - **ordered sequence of values**, where the **same value may occur more than once**.
Lists are a basic example of containers.

* Stack - is an abstract data type that serves as a **collection** of elements,
with two principal operations: push and pop. (**LIFO**).

* Queue - is a particular kind of abstract data type or **collection** in which the entities in the collection are kept in order.
(**FIFO**).

* Double-ended queue - (often abbreviated to deque) is an abstract data type that generalizes a queue,
for which **elements can be added to or removed from either the front (head) or back (tail)**.
(also often called a head-tail linked list).

* Priority queue - is an abstract data type which is like a regular queue or stack data structure,
but where additionally **each element has a "priority"** associated with it.

* Double-ended priority queue - is a data structure similar to a priority queue,
but allows for **efficient removal of both the maximum and minimum**.

* Set - is an abstract data type that can **store certain values**,
**without** any particular **order**, and **no repeated values**.

* Multiset (or bag) - is similar to a set but allows repeated ("equal") values (**duplicates**).

* Disjoint-set (union–find or merge–find) - is a data structure
that keeps track of a set of elements partitioned into a number of disjoint (nonoverlapping) subsets.

#### Arrays

* Bit array (bitmap, bitset, bit string, or bit vector) - is an array data structure that compactly **stores bits**.
Is effective at exploiting bit-level parallelism in hardware to perform operations quickly.

* Circular buffer (circular queue, cyclic buffer or ring buffer) -  is a data structure
that uses a single, fixed-size buffer as if it were **connected end-to-end**.

* Dynamic array (growable array, resizable array, dynamic table, mutable array, or array list) -
is a **random access, variable-size** list data structure that allows elements to be added or removed.

* Hash table (hash map) - is a data structure
used to implement an associative array, a structure that can **map keys to values**. A hash table uses a hash function
to compute an index into an array of buckets or slots, from which the desired value can be found.
The average cost for each lookup is independent of the number of elements stored in the table.
Many hash table designs also allow arbitrary insertions and deletions of key-value pairs.
Widely used in many kinds of computer software,
particularly for associative arrays, database indexing, caches, and sets.

* Hashed array tree - is a dynamic array data-structure
maintaining an array of **separate memory fragments (or "leaves")** to store the data elements,
unlike simple dynamic arrays which maintain their data in one contiguous memory area.

* Sparse array - is an array in which most of the elements have the **default value** (usually 0 or null).

#### Linked

* Linked list - is a linear collection of data elements,
called nodes pointing to the next node by means of a pointer.
It is a data structure consisting of a group of nodes which together represent a sequence.
Under the simplest form, each node is composed of data and a **reference
to the next node** in the sequence.
The principal benefit - is that the list elements can easily be inserted or removed
without reallocation or reorganization of the entire structure
because the data items need not be stored contiguously in memory or on disk.

* Doubly linked list - each node contains two links,
that are **references to the previous and to the next node** in the sequence of nodes.

* Association list (alist) - is a linked list
in which each list **element (or node) comprises a key and a value**.
The association list is said to associate the value with the key.

* Skip list - is a data structure that allows **fast search** within an ordered sequence of elements,
by maintaining a linked hierarchy of subsequences.

* Unrolled linked list - is a variation on the linked list which **stores multiple elements in each node**.
It is related to the B-tree.

* XOR linked list - it takes advantage of the bitwise XOR operation
to decrease storage requirements for doubly linked lists.

#### Trees

* Binary search trees (ordered, sorted binary trees) - particular type of containers,
data structures that store "items" in memory.
They allow fast lookup, addition and removal of items,
and can be used to implement either dynamic sets of items,
or lookup tables that allow finding an item by its keys.
Keep their keys in sorted order, so that lookup and other operations can **use the principle of binary search**.

* Self-balancing binary search tree (height-balanced) - is any node-based binary search tree
that automatically **keeps its height small** in the face of arbitrary item insertions and deletions.

* B-tree - is a self-balancing tree data structure
that keeps data sorted and allows searches,
sequential access, insertions, and deletions in logarithmic time.
The B-tree is a generalization of a binary search tree in that a node **can have more than two children**.

* Red–black tree - is a kind of self-balancing binary search tree.
Each node of the binary tree has an extra bit,
and that bit is often interpreted as the color (red or black) of the node.
These color bits are used to ensure the tree remains approximately balanced during insertions and deletions.

* 2–3 tree:
2-node if it has one data element and two children.
3-node if it has two data elements and three children.

* AA tree (named for Arne Andersson, their inventor) - variation of the red-black tree,
is a form of binary search tree which supports efficient addition and deletion of entries.
Unlike red-black trees, red nodes on an AA tree can only be added as a right subchild.

* AVL tree - the heights of the two child subtrees of any node differ by at most one.
Are faster than red–black trees because they are more rigidly balanced.
Similar to red–black trees, AVL trees are height-balanced.

* Splay tree -  is a self-adjusting binary search tree
with the additional property that recently accessed elements are quick to access again.
All normal operations on a binary search tree are combined with one basic operation, called splaying.

* Heap - can be classified further as either a "max heap" or a "min heap".
In a max heap, the keys of parent nodes are always greater than or equal
to those of the children and the highest key is in the root node.
In a min heap, the keys of parent nodes are less than or equal.

* Binary heap - is a heap data structure created using a binary tree.

* Binomial heap - is a heap similar to a binary heap
but also supports quick merging of two heaps.

* Fibonacci heap - is a data structure for priority queue operations,
consisting of a collection of heap-ordered trees.

* R-tree - tree data structures used for spatial access methods.

* R*-trees - have slightly higher construction cost than standard R-trees,
as the data may need to be reinserted;
but the resulting tree will usually have a better query performance.

* R+ tree - is a method for looking up data using a location, often (x, y) coordinates,
and often for locations on the surface of the earth.

* Hilbert R-tree - an extension to B+-tree for multidimensional objects.

* Trie (digital or radix or prefix tree) - is an ordered tree data structure
that is used to store a dynamic set or associative array where the keys are usually strings.

* Hash tree - is a persistent data structure that can be used to implement sets and maps,
intended to replace hash tables in purely functional programming.

* Binary decision diagram (BDD) -  is a data structure
that is used to represent a Boolean function.
On a more abstract level, BDDs can be considered
as a compressed representation of sets or relations.

* Directed acyclic graph -

http://en.wikipedia.org/wiki/Template:Data_structures
