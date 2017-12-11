/**
 * Find most frequent element in array.
 *
 * @todo: Finish it.
 *
 * @param {array} a Array with numbers.
 * @return {number} Index in array to most frequent element in array.
 */
function  f(a) {
  var arr = [...a];
  var len = arr.length;
  for (var i = 0; i < len; i++) {
    arr[arr[i]%len] += len;
  }
  var max = arr[0];
  var result = 0;
  for (var i = 1; i < len; i++) {
    if (arr[i] > max) {
      max = arr[i];
      result = i;
    }
  }
  for (var i = 0; i < a.length; i++) {
    if (a[i] == result) {
      return i;
    }
  }
  return -1;
}

// console.log(f([3, 4, 3, 2, 3, -1, 3, 3]));
// console.log(f([]));
console.log(f([2, 1, 4, 7, 4, 8, 3, 6, 4, 7]));
