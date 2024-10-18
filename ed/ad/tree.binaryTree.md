Binary Tree
-

**Tree** - connected graph with N nodes & N-1 edges.

Tree center - middle vertex (1 or 2) in longest path.
Isomorphic Tree - special case of graph isomorphism.

Storing undirected trees:
* Edge list: [(0,1), (1,2), (0,3)].
* Adjacency list: 0 -> [1,3], 1 -> [2].
* Adjacency matrix: axis x - nodes values; axis y - nodes values, cells - edge between nodes.

**Rooted tree** - tree with designated root node.

Storing rooted tree:
* Flattened array: [0,1,3,2].

**Binary tree** - data structure where each node has at most 2 children.

Height of BT = 1 + numbers of edges on the longest path from root to leaf.
Diameter of a BT - longest path of tree between 2 leafs.
Delete from BT = Postorder + function free().

By knowing that node may have at most 2 childs - BT may be stored as array.

When BT saved in array, elements can be found by:
````
i                        - element index
2*i          2*i+1       - left child
2*i+1        2*i+2       - right child
floor(i/2)  floor(i-1/2) - parent
````

Full BT - every non-leaf node has two children,
no space to add element into tree without affecting tree's height.

Complete BT - if store BST in array, there must NOT be missing elements in array (`[a,b,-,d]`).

BFS/DFS (Breadth/Depth First Search) in binary tree: @see: graph.md
Level-Order Traversal aka BFS.
Preorder Traversal aka DFS.

Tree rotation - moves one node up in BT and one node down, to decrease its height,
by moving smaller subtrees down and larger subtrees up.

````
rotate right ->

    50         17
   / \        / \
  17  76     9  50
 / \            / \
9  23          23  76

       <- rotate left
````

**Binary Search Tree** (BST) (ordered, sorted binary trees) - can **use the principle of binary search**,
particular type of containers, data structures that store items in memory.
They allow fast lookup, addition and removal of items,
and can be used to implement either dynamic sets of items,
or lookup tables that allow finding an item by its keys.
Keep their keys in sorted order, so that lookup and other operations can **use the principle of binary search**.

BST invariant: `x.left.value <= x.value <= x.right.value`.
BST uniqueness: `x.left.value < x.value < x.right.value`.

BST is balanced if depth of two subtrees of every node never differs by more than 1.
Balanced BST - BST that has minimum possible height.

Get BST height: recursive call height for left & right; return max(left height, right height)+1.

Insert into BST: `val < left.val -> insert into left node, else insert into right node`.

Delete node from BST:
1. delete leaf node - just delete it.
2. delete node with 1 child - replace node with it's child.
3. delete node with 2 childs - replace with node wich is minimum in right child.

Valid BST:
1. Left subtree contains only nodes with keys less than the node's key.
2. Right subtree contains only nodes with keys greater than the node's key.
3. Both left and right subtrees also BST.

Print Root to Leaf Path with Given sum (K-Sum paths):
````
1. push root value into stack
2. -> go into left child push value into stack and calculate stack sum
  -> if it's:
    node - go to step 2
    leaf - pop value from stack and go into parent right child and go to step 2
````

Find inorder predecessor:
If left child is present - got to left child and go to most right,
otherwise - search from where we take the last right turn.

Find inorder successor:
If right child is present - got to right child and go to most left,
otherwise - search from where we take the last left turn.

Preorder traversal (DFS): print node; go to left child; go to right child;
Postorder traversal (DFS): go to left child; go to right child; print node;
Inorder traversal (DFS): go to left child; print node; go to right child;
Level traversal (BFS);

Vertical order traversal:
1. put in queue root, root height = 0.
2. deque, left child height = parent height -1, right child height = parent height +1.
3. enqueue left child & right child.

Spiral (zig-zag) traversal of a binary tree:
1. add root to Stack1.
2. pop all from Stack1 and push left child into Stack2 and right child into Stack2.
3. pop all from Stack2 and push right child into Stack1 and left child into Stack1.
4. go to step 2.

Diagonal distance:
````
root       = 0
left child = parent d - 1
left child = parent d
````

Number of Binary Search Trees possible with N nodes:
For example, `for: [5, 6] result: [5, 6], [6, 5] (all permutations)`.

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

**Threaded binary tree** (TBT) - every node have value and left pointer and right pointer.
In case we have left or right pointer empty - we can fill it with link to inorder predecessor/successor,
so it become a TBT.
To differentiate is it a pointer to child or to predecessor/successor TBT has left and right flags.
