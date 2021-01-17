Tree
-

Balancing tree:
* Right rotation (when tree is left heavy)
* Left rotation (when tree is right heavy)
* Right-Left rotation
* Left-Right rotation

````
Update segment tree            - O(log n).
Update range tree              - O(log n).
Update range array (tree)      - O(n).
Select from range array (tree) - O(1).
````
<br>Segment tree - better for many updates.
<br>Range array (tree) - better for many reads.
<br>Diameter of a Binary Tree - longest path of tree between 2 leafs.
<br>2 trees are isomorphic - if subtrees are isomorphic to each other.
<br>Height of BT = 1 + numbers of edges on the longest path from root to leaf.
<br>Delete from Binary Tree = Postorder + function free()

To make Minimum spanning tree use:
* Prim's Algorithm - like BFS but pick chipest node on each iteration.
* Kruskal's Algorithm - connect to endes while them in different separated trees, because cycles not allowed.
* Boruvka's Algorithm
* Reverse-Delete Algorithm

## Trees

* AST - Abstract Syntax Tree (used by webpack, babel, etc). https://astexplorer.net/

* `Binary search trees` (ordered, sorted binary trees) - particular type of containers,
data structures that store items in memory.
They allow fast lookup, addition and removal of items,
and can be used to implement either dynamic sets of items,
or lookup tables that allow finding an item by its keys.
Keep their keys in sorted order, so that lookup and other operations can **use the principle of binary search**.

* `Self-balancing binary search tree` (height-balanced) - is any node-based binary search tree
that automatically **keeps its height small** in the face of arbitrary item insertions and deletions.
Search performance = O(log n)

* `B-tree` - self-balancing, self managed, multi-level binary search tree.

B-tree is a self-balancing tree data structure
that keeps data sorted and allows searches,
sequential access, insertions, and deletions in logarithmic time.
The B-tree is a generalization of a binary search tree in that a node **can have more than two children**.

B-tree is also self managed multi-level index,
it means when data in index grows, B-tree creates index in front of index, and so on,
and when data deletes, B-tree deletes index.

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

* `R-tree` - tree data structures **used for spatial access methods**.

* `R*-trees` - have slightly higher construction cost than standard R-trees,
as the data may need to be reinserted;
but the resulting tree will usually have a better query performance.

* R+ tree - is a method for looking up data using a location, often (x, y) coordinates,
and often for locations on the surface of the earth.

* Hilbert R-tree - an extension to B+-tree for multidimensional objects.

* Hash tree - is a persistent data structure that can be **used to implement sets and maps**,
intended to replace hash tables in purely functional programming.

* Log-structured merge-tree (LSM) - tree with performance characteristics that make it attractive
for providing indexed access to files with high insert volume, such as transactional log data.
