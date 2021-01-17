Sorts
-

Performance:
* number of comparisons
* number of swaps

Performance also depends on size and data structure presented in array.

### [Bubble sort](https://upload.wikimedia.org/wikipedia/commons/c/c8/Bubble-sort-example-300px.gif)

Worst case = **O(n^2)**; Average case = O(n^2); Best case = O(n)

Good for small arrays. Bad for big arrays.
No extra memory allocations needed for bubble sort!

### Selection sort

Worst case = **O(n^2)**; Average case = O(n^2); Best case = O(n^2)
Space required = O(n)

Good for small arrays. Bad for big arrays.
In practice it better than bubble sort but worst than insertion sort.
Don\'t do many swaps but do lot of comparisons.

### [Insertion sort](https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif)

Worst case = **O(n^2)**; Average case = O(n^2); Best case = O(n)

Good for small arrays. Bad for big arrays.
No extra memory allocations needed for insertion sort!

### Quick sort

reqursive

Worst case = **O(n^2)**; Average case = O(n log n); Best case = O(n log n)
Space required = O(n)

### [Merge sort](https://upload.wikimedia.org/wikipedia/commons/c/cc/Merge-sort-example-300px.gif)

reqursive

Worst case = **O(n log n)** (less than O(n^2)); Average case = O(n log n); Best case = O(n log n)
Space required = O(n)

Provides ability to perform sub-sorts in parallel.
Predicteble algorithm because only size of array influence performance.

### Heap sort

Worst case = **O(n log n)**
