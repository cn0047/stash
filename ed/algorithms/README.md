Algorithms
-

## Sorts

Performance:

* number of comparisons
* swaps

Performance also depends on size and data structure presented in array.

### Bubble sort

[Example](https://upload.wikimedia.org/wikipedia/commons/c/c8/Bubble-sort-example-300px.gif)

Worst case = O(n^2)
Average case = O(n^2)
Best case = O(n)

Good for small arrays. Bad for big arrays.
No extra memmory allocations needed for bubble sort!

### Insertion sort

[Example](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

Worst case = O(n^2)
Average case = O(n^2)
Best case = O(n)

Good for small arrays. Bad for big arrays.
No extra memmory allocations needed for bubble sort!

### Selection sort

Worst case = O(n^2)
Average case = O(n^2)
Best case = O(n^2)
Space required = O(n)

Good for small arrays. Bad for big arrays.
In practice it better than bubble sort but worst than insertion sort.
Don't do many swaps but do lot of comparisons.

### Merge sort

[Example](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

Worst case = O(n log n) # less than O(n^2)
Average case = O(n log n)
Best case = O(n log n)
Space required = O(n)

Provides ability to perform sub-sorts in parallel.
Predicteble algorithm because only size of array influence performance.

### Quick sort

Worst case = O(n^2)
Average case = O(n log n)
Best case = O(n log n)
Space required = O(n)

## Tree

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

## String search

### Naive search algorithm

Performance = O(n+m) where n is length of str to search and m length of str to find
Worst performance = O(n*m)

### Boyer-More-Horspool search
