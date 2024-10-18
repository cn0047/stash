Tree
-

<br>Range array (tree) - better for many reads.
<br>2 trees are isomorphic - if subtrees are isomorphic to each other.

````
Update segment tree            - O(log n).
Update range tree              - O(log n).
Update range array (tree)      - O(n).
Select from range array (tree) - O(1).
````

To make Minimum spanning tree use:
* Prim's Algorithm - like BFS but pick chipest node on each iteration.
* Kruskal's Algorithm - connect to endes while them in different separated trees, because cycles not allowed.
* Boruvka's Algorithm
* Reverse-Delete Algorithm

**AST** - Abstract Syntax Tree (used by webpack, babel, etc).
[explorer](https://astexplorer.net/)

**2–3 tree**:
2-node if it has one data element and two children.
3-node if it has two data elements and three children.

**AA tree** (named for inventor Arne Andersson) - variation of the red-black tree,
is a form of binary search tree which supports efficient addition and deletion of entries.
**Unlike red-black** trees, **red nodes** on an AA tree **can only be added as a right subchild**.

**Splay tree** - is a self-adjusting **BST with additional property - recently accessed elements**.
Splay tree contains same basic BST operations (insert, delete, search) and after every operation
there is additional operation - splaying.
After each operation tree must be rebalanced and target node of operation must become root.

Types of rotations:
* Zig rotation.
* Zag rotation.
* Zig-Zig rotation.
* Zag-Zag rotation.
* Zig-Zag rotation.
* Zag-Zig rotation.

Applications: caching, DB indexing, file systems, data compression, text processing, graph algorithms, online gaming.

**R-tree** - tree data structures **used for spatial access methods**.

**R*-tree** - has slightly higher construction cost than standard R-tree,
as the data may need to be reinserted;
but the resulting tree will usually have a better query performance.

**R+ tree** - is a method for looking up data using a location, often (x, y) coordinates,
and often for locations on the surface of the earth.

**Hilbert R-tree** - an extension to B+-tree for multidimensional objects.

**Hash tree** - is a persistent data structure that can be **used to implement sets and maps**,
intended to replace hash tables in purely functional programming.

**Log-structured merge-tree** (LSM) - tree with performance characteristics that make it attractive
for providing indexed access to files with high insert volume, such as transactional log data.

**AVL tree** - (since 1962) BST where **heights of the two child subtrees of any node differ by at most one**.
Faster than red–black trees because they are more rigidly balanced.
Similar to red–black trees, AVL trees are height-balanced.

Balance factor: +1, 0, -1
Rotations: LL, RR, LR, RL.

Height: `h = log N`.
Worst time complexity for search: `h = O(h) = log N`.

**Radix tree** - each node that is the only child is merged with its parent.
This makes radix trees much more efficient for small sets
(especially if the strings are long) and for sets of strings
that share long prefixes.

[**Trie**](https://monosnap.com/file/1sTq5fwBiVEvPWakJiGusGK7foFmji)
(digital or radix or prefix tree) - is an ordered tree data structure
that is used to **store dynamic set or associative array** where the keys are usually strings.

Root node can have any number of nodes.
Root does not contain value except the pointers to its child nodes.

Types of trie data structures:
* Standard tries.
* Compressed tries.
* Suffix tries.

**Heap** - complete BST.
````
find max               - O(1)
insert                 - O(log n)
delete                 - O(log n)
create heap from array - O(n log n)
# heapify              - O(n)
````

In a max heap, the keys of parent nodes are always greater than or equal to those of the children
and the highest key is in the root node.
In a min heap, the keys of parent nodes are less than or equal.

Binary heap - is a heap data structure created using a binary tree.
Binomial heap - is a heap similar to a binary heap
but also supports **quick merging of two heaps**.
Fibonacci heap - is a data structure for **priority queue operations,
consisting of a collection of heap-ordered trees**.

Only root element can be deleted from heap (not any arbitrary element).

**Segment tree** - binary tree whose nodes store information about segments.
Good to solve range min/max/sum/geographic information systems/computational geometry
queries and range updates in O(log n) time.
Three operations: build tree, update tree, query tree.
