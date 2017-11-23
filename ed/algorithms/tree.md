Tree
-

Operations with sets ([1, 2, 3] [2, 3, 4]):

* union
* intersection
* diff
* symmetric diff

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

## AVL

Balance factor: +1, 0, -1

LL | RR rotations.

Height h = log N
Worst time complexity for search = h = log N

## Binary Search

#### Find inorder predecessor

If left child is present - got to left child and go to [most right](http://prntscr.com/hdpp78),
otherwise - search from where we take the [last right turn](http://i.prntscr.com/N07a6FMpQxy0ho1XoQ0RdQ.png).

#### Find inorder successor

If right child is present - got to right child and go to [most left](http://prntscr.com/hdpsl5),
otherwise - search from where we take the [last left turn](http://prntscr.com/hdptzo).

#### Number of Binary Search Trees possible with N nodes

For example, for: [5, 6] Result: [5, 6], [6, 5]

#### TBT (Threaded Binary Tree)

Every node have value and left pointer and right pointer.
In case we have left or right pointer empty - we can fill it with link to inorder predecessor/successor,
so it become a TBT.
To differentiate is it a pointer to child or to predecessor/successor TBT has left and right flags.

#### Postorder Traversal (Shortcut Trick)

1. go to left child
2. go to right child
3. print node
