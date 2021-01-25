// Insertion sort.
// @see https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif
// Worst case = O(n^2).
// Average case = O(n^2).
// Best case = O(n).
// Good for small arrays. Bad for big arrays.
// No extra memory allocations needed for insertion sort!
//
// The plan:
// just like sorting bunch of cards in hands.
function insertionSort(arr) {
  for (let i = 1; i < arr.length; i++) {
    let j = i - 1;
    do {
      let v = arr[j + 1];
      if (v < arr[j]) {
        arr[j + 1] = arr[j];
        arr[j] = v;
      }
      j--;
    } while (j >= 0);
  }
  return arr;
}

console.log(insertionSort([5, 3, 1, 2, 4]));
console.log(insertionSort([2, 4, 6, 8, 3]));
console.log(insertionSort([2, 4, 6, 8, 9]));
console.log(insertionSort([2, 4, 6, 8, 0]));
