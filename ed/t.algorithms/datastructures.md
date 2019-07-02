Data structures
-

## Abstract

* `Container` - is a class, a data structure, or an abstract data type (ADT)
whose instances are collections of other objects. In other words, container store **objects in an organized way**.

* `Collection` - is a grouping of some variable number of data items (possibly zero)
that have some shared significance to the problem being solved (Generally, the data items will be of the **same type**).

* `Associative array`, map, symbol table, or dictionary - is an abstract data type
composed of a collection of **(key, value) pairs**, such that each possible **key appears just once** in the collection.

* `Multimap` generalizes an associative array by allowing **multiple values to be associated with a single key**.

* `Bidirectional map` - is a related abstract data type in which the **bindings operate in both directions**:
each value must be associated with a unique key,
and a second lookup operation takes a value as argument and looks up the key associated with that value.

* `List` - **ordered sequence of values**, where the **same value may occur more than once**.
Lists are a basic example of containers.

* `Stack` - is an abstract data type that serves as a **collection** of elements,
with two principal operations: push and pop. (**LIFO**).

* `Queue` - is a particular kind of abstract data type or **collection** in which the entities in the collection are kept in order.
(**FIFO**).

* `Double-ended queue` - (often abbreviated to deque) is an abstract data type that generalizes a queue,
for which **elements can be added to or removed from either the front (head) or back (tail)**.
(also often called a head-tail linked list).

* `Priority queue` - is an abstract data type which is like a regular queue or stack data structure,
but where additionally **each element has a "priority"** associated with it.

* `Double-ended priority queue` - is a data structure similar to a priority queue,
but allows for **efficient removal of both the maximum and minimum**.

* `Set` - is an abstract data type that can **store certain values**,
**without** any particular **order**, and **no repeated values**.

* `Multiset` (or bag) - is similar to a set but allows repeated ("equal") values (**duplicates**).

* Disjoint-set (union–find or merge–find) - is a data structure
that keeps track of a set of elements partitioned into a number of disjoint (nonoverlapping) subsets.

## Linked

* `Linked list` - is a linear **collection** of data elements,
called nodes pointing to the next node.
It is a data structure consisting of a group of nodes which together represent a sequence.
Under the simplest form, each node is composed of data and a **reference
to the next node** in the sequence.
The principal benefit - is that the list elements can easily be inserted or removed
without reallocation or reorganization of the entire structure
because the data items need not be stored contiguously in memory or on disk.
<br>Disadvantage: slow to get the element; slow append to end;
Advantage: fast insert/delete into/from head;

* `Doubly linked list` - each node contains two links,
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

http://en.wikipedia.org/wiki/Template:Data_structures
