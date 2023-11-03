Tree
-

<br>Segment tree - better for many updates.
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

**Threaded binary tree** (TBT) - every node have value and left pointer and right pointer.
In case we have left or right pointer empty - we can fill it with link to inorder predecessor/successor,
so it become a TBT.
To differentiate is it a pointer to child or to predecessor/successor TBT has left and right flags.

**Self-balancing binary search tree** (height-balanced) - is any node-based binary search tree
that automatically **keeps its height small** in the face of arbitrary item insertions and deletions.
Search performance = `O(log n)`.

Balancing tree:
* Right rotation (when tree is left heavy).
* Left rotation (when tree is right heavy).
* Right-Left rotation.
* Left-Right rotation.

**B-tree** - self-balancing, self managed, multi-level binary search tree.

B-tree is a self-balancing tree data structure
that keeps data sorted and allows searches,
sequential access, insertions, and deletions in logarithmic time.
The B-tree is a generalization of a binary search tree in that a node **can have more than two children**.

B-tree is also self managed multi-level index,
it means when data in index grows, B-tree creates index in front of index, and so on,
and when data deletes, B-tree deletes index.

**AST** - Abstract Syntax Tree (used by webpack, babel, etc).
[explorer](https://astexplorer.net/)

**Red–black tree** - is a kind of self-balancing binary search tree.
**Each node** of the binary tree **has an extra bit**,
and that bit is often interpreted as the color (red or black) of the node.
These color bits are used **to ensure the tree remains approximately balanced** during insertions and deletions.

**2–3 tree**:
2-node if it has one data element and two children.
3-node if it has two data elements and three children.

**AA tree** (named for inventor Arne Andersson) - variation of the red-black tree,
is a form of binary search tree which supports efficient addition and deletion of entries.
**Unlike red-black** trees, **red nodes** on an AA tree **can only be added as a right subchild**.

**Splay tree** - is a self-adjusting binary search **tree
with the additional property that recently accessed elements** are quick to access again.
All normal operations on a binary search tree are combined with one basic operation, called splaying.

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

**AVL tree** - (since 1962) the **heights of the two child subtrees of any node differ by at most one**.
Faster than red–black trees because they are more rigidly balanced.
Similar to red–black trees, AVL trees are height-balanced.

Balance factor: +1, 0, -1

LL | RR rotations.

Height: `h = log N`.
Worst time complexity for search: `h = O(h) = log N`.

**Radix tree** - each node that is the only child is merged with its parent.
This makes radix trees much more efficient for small sets
(especially if the strings are long) and for sets of strings
that share long prefixes.

[Trie](https://monosnap.com/file/1sTq5fwBiVEvPWakJiGusGK7foFmji)
(digital or radix or prefix tree) - is an ordered tree data structure
that is used to **store a dynamic set or associative array** where the keys are usually strings.

**Heap** - complete BST.
````
find max               - O(1)
insert                 - O(log n)
delete                 - O(log n)
create heap from array - O(n log n)
# heapify              - O(n)
````

**In a max heap, the keys of parent nodes are always greater than or equal to those of the children**
and the highest key is in the root node.
In a min heap, the keys of parent nodes are less than or equal.

Binary heap - is a heap data structure created using a binary tree.
Binomial heap - is a heap similar to a binary heap
but also supports **quick merging of two heaps**.
Fibonacci heap - is a data structure for **priority queue operations,
consisting of a collection of heap-ordered trees**.

Only root element can be deleted from heap (not any arbitrary element).
