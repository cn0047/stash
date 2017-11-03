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

#### AVL

Balance factor: +1, 0, -1

LL | RR rotations.

Height h = log N
Worst time complexity for search = h = log N

#### Binary Search

Inorder predecessor:
If left element is present - got to left element and go to most right,
otherwise - search from where we take the last right turn.

Inorder successor:
