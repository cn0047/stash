Tree
-

Balancing tree:

* Right rotation (when tree is left heavy)
* Left rotation (when tree is right heavy)
* Right-Left rotation
* Left-Right rotation

Update segment tree - O(log N).
Update range tree - O(log N).

Select from range array (tree) - O(1).
Update range array (tree) - O(N).

Segment tree - better for many updates.
Range array (tree) - better for many reads.

Diameter of a Binary Tree - longest path of tree between 2 leafs.

2 trees are isomorphic - if subtrees are isomorphic to each other.

Height of BT = 1 + numbers of edges on the longest path from root to leaf.

Delete a Binary Tree = Postorder + function free()

To make Minimum spanning tree use:
* Prim's Algorithm - like BFS but pick chipest node on each iteration.
* Kruskal's Algorithm - connect to endes while them in different separated trees,
  because cycles not allowed.
* Boruvka's Algorithm
* Reverse-Delete Algorithm

## Trees

AST - Abstract Syntax Tree (used by webpack, babel, etc). https://astexplorer.net/

* `Binary search trees` (ordered, sorted binary trees) - particular type of containers,
data structures that store "items" in memory.
They allow fast lookup, addition and removal of items,
and can be used to implement either dynamic sets of items,
or lookup tables that allow finding an item by its keys.
Keep their keys in sorted order, so that lookup and other operations can **use the principle of binary search**.

    * Pre-order traversal - process inself -> visit left node -> visit right node; repeat recursive for each node.
    * In-order traversal (sort order) -.visit left node -> process inself -> visit right node; repeat recursive for each node.
    * Post-order traversal - visit left node -> visit right node -> process inself; repeat recursive for each node.

* `Self-balancing binary search tree` (height-balanced) - is any node-based binary search tree
that automatically **keeps its height small** in the face of arbitrary item insertions and deletions.
Search performance = O(log n)

* `B-tree` - is a self-balancing tree data structure
that keeps data sorted and allows searches,
sequential access, insertions, and deletions in logarithmic time.
The B-tree is a generalization of a binary search tree in that a node **can have more than two children**.

* Red–black tree - is a kind of self-balancing binary search tree.
**Each node** of the binary tree **has an extra bit**,
and that bit is often interpreted as the color (red or black) of the node.
These color bits are used **to ensure the tree remains approximately balanced** during insertions and deletions.

* 2–3 tree:
2-node if it has one data element and two children.
3-node if it has two data elements and three children.

* AA tree (named for Arne Andersson, their inventor) - variation of the red-black tree,
is a form of binary search tree which supports efficient addition and deletion of entries.
**Unlike red-black** trees, **red nodes** on an AA tree **can only be added as a right subchild**.

* Splay tree - is a self-adjusting binary search **tree
with the additional property that recently accessed elements** are quick to access again.
All normal operations on a binary search tree are combined with one basic operation, called splaying.

* `Heap` - can be classified further as either a "max heap" or a "min heap".
**In a max heap, the keys of parent nodes are always greater than or equal
to those of the children** and the highest key is in the root node.
In a min heap, the keys of parent nodes are less than or equal.

O(1) - find max
O(log n) - insert
O(log n) - delete

* Binary heap - is a heap data structure created using a binary tree.

* Binomial heap - is a heap similar to a binary heap
but also supports **quick merging of two heaps**.

* Fibonacci heap - is a data structure for **priority queue operations,
consisting of a collection of heap-ordered trees**.

* `R-tree` - tree data structures **used for spatial access methods**.

* R*-trees - have slightly higher construction cost than standard R-trees,
as the data may need to be reinserted;
but the resulting tree will usually have a better query performance.

* R+ tree - is a method for looking up data using a location, often (x, y) coordinates,
and often for locations on the surface of the earth.

* Hilbert R-tree - an extension to B+-tree for multidimensional objects.

* Hash tree - is a persistent data structure that can be **used to implement sets and maps**,
intended to replace hash tables in purely functional programming.

## AVL

AVL tree - (since 1962) the **heights of the two child subtrees of any node differ by at most one**.
Are faster than red–black trees because they are more rigidly balanced.
Similar to red–black trees, AVL trees are height-balanced.

Balance factor: +1, 0, -1

LL | RR rotations.

Height h = log N
Worst time complexity for search = h = O(h) = log N

## Radix tree

* [Trie](https://monosnap.com/file/1sTq5fwBiVEvPWakJiGusGK7foFmji) (digital or radix or prefix tree) - is an ordered tree data structure
that is used to **store a dynamic set or associative array** where the keys are usually strings.

Radix tree - each node that is the only child is merged with its parent.
This makes radix trees much more efficient for small sets
(especially if the strings are long) and for sets of strings
that share long prefixes.
