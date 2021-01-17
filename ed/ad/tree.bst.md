Binary Search Tree (BST)
-

Level-Order Traversal aka BFS.

Full BT - every node other than the leaves has two children,
no space to add element into tree without affecting tree's height.

Complete BT - if store BST in array, there must NOT be missing elements in array (`[a,b,-,d]`).

#### Heap

Heap - complete BST.

````
find max               - O(1)
insert                 - O(log n)
delete                 - O(log n)
create heap from array - O(n log n)
heapify                - O(n)
````

**In a max heap, the keys of parent nodes are always greater than or equal to those of the children**
and the highest key is in the root node.
In a min heap, the keys of parent nodes are less than or equal.

Binary heap - is a heap data structure created using a binary tree.
Binomial heap - is a heap similar to a binary heap
but also supports **quick merging of two heaps**.
Fibonacci heap - is a data structure for **priority queue operations,
consisting of a collection of heap-ordered trees**.

#### Find inorder predecessor

If left child is present - got to left child and go to [most right](http://prntscr.com/hdpp78),
otherwise - search from where we take the [last right turn](http://i.prntscr.com/N07a6FMpQxy0ho1XoQ0RdQ.png).

#### Find inorder successor

If right child is present - got to right child and go to [most left](http://prntscr.com/hdpsl5),
otherwise - search from where we take the [last left turn](http://prntscr.com/hdptzo).

#### Number of Binary Search Trees possible with N nodes

For example, `for: [5, 6] result: [5, 6], [6, 5] (all permutations)`.

#### Threaded binary tree (TBT)

Every node have value and left pointer and right pointer.
In case we have left or right pointer empty - we can fill it with link to inorder predecessor/successor,
so it become a TBT.
To differentiate is it a pointer to child or to predecessor/successor TBT has left and right flags.

#### Preorder Traversal (Shortcut Trick)

1. print node
2. go to left child
3. go to right child

#### Postorder Traversal (Shortcut Trick)

1. go to left child
2. go to right child
3. print node

#### Inorder Traversal (Shortcut Trick)

1. go to left child
2. print node
3. go to right child

### BFS/DFS (Breadth/Depth First Search) in binary tree

@look: graph.md

### Delete a node from Binary Search Tree

Delete leaf node - just delete it.
<br>Delete node with 1 child - replace node with it's child.
<br>Delete node with 2 childs - replace with node wich is minimum in right child.

### Print Root to Leaf Path with Given sum (K-Sum paths)

````
1. push root value into stack
2. -> go into left child push value into stack and calculate stack sum
  -> if it's:
    node - go to step 2
    leaf - pop value from stack and go into parent right child and go to step 2
````

### Spiral (zig-zag) traversal of a binary tree

1. add root to Stack1
2. pop all from Stack1 and push left child into Stack2 and right child into Stack2
3. pop all from Stack2 and push right child into Stack1 and left child into Stack1
4. go to step 2

### Diagonal distance

````
root       = 0
left child = parent d - 1
left child = parent d
````
