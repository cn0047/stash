Red-black Tree
-

````
      B100
     /    \
  R60      B145
 /   \     /   \
B21  B76  R110  R150
  \
  R32

     B4
    /   \
  R2     R6
 /  \    / \
B1  B3  B5  B8
            / \
           R7  R9

     B6
    /   \
  R3     R8
 /  \    / \
B1  B5  B7  B9
             \
              R10
````

Search: O(log n).
Insert: O(log n).
Delete: O(log n).

**Red-black tree** - is a kind of self-balancing binary search tree.
**Each node** of the binary tree **has an extra bit**,
and that bit is often interpreted as the color (red or black) of the node,
these color bits are used **to ensure the tree remains approximately balanced** during insertions and deletions.

Parent nodes are black, children nodes are red.
The leaf nodes (NIL) do not contain keys or data.

Properties:
1. Every node is either red or black. Root is alway black.
2. All NIL nodes are considered black.
3. Red node does not have red child. Black node may have black child.
4. Every path from given node to any of its descendant NIL nodes goes through the same number of black nodes.
5. If node has exactly one child, it must be a red child.

#### Search

Searching is similar to searching in standard BST.

#### Insert

New inserted node is always red.

1. BST insert.
2. Fix violations (if parent of the new node is black - ok, if red - fix).

Fix cases:
1. Uncle is red: Recolor parent and uncle to black, grandparent to red.
2. Uncle is black:
  If node is right child - perform left rotation on the parent.
  If node is left child - perform right rotation on the grandparent and recolor appropriately.

Fix:
1. Recoloring and propagating upwards:
  If parent and uncle of the new node are both red - recolor parent and uncle to black, grandparent to red.
  Recursively apply the fix-up to the grandparent.
2. Rotation and recoloring:
  If the new node's uncle is black and the new node is the right child of left child (or left of right) - perform
  rotation to move the new node up and align it.
  If the new node's uncle is black and the new node is the left child of left child (or right of right) - perform
  rotation and recolor the parent and grandparent to fix the violation.

#### Delete

1. BST delete.
2. Fix violations.

Fix cases:
1. Sibling is red: Rotate parent and recolor sibling and parent.
2. Sibling is black & sibling's children are black: Recolor sibling and propagate the double black upwards.
3. Sibling is black & at least one of the sibling's children is red:
  If the sibling's far child is red: perform rotation on the parent and sibling, and recolor appropriately.
  If the sibling's near child is red: rotate the sibling and its child, then handle as above.

Fix:
1. Sibling is red: Recolor the sibling and parent, and perform rotation.
2. Sibling is Black with Black Children: Recolor the sibling to red and move the problem up to parent.
3. Sibling is Black with at least one Red Child: Rotate and recolor to fix the double-black issue.

#### Rotations

Rotations in Red-Black trees are typically performed during insertions and deletions.
