/**
 * Gets MAX possible triangle perimeter from array with numbers.
 *
 * @param {array} arr Input array with numbers.
 * @return {number} Max possible triangle perimeter.
 */
function f1(arr) {
    if (arr.length === 3) return -1;
    var max = 0;
    while (arr.length > 0) {
      var v1 = arr.shift();
      for (var i = 0; i < arr.length - 1; i++) {
        for (var j = 1; j < arr.length; j++) {
          if (i === j) continue;
          var v2 = arr[i];
          var v3 = arr[j];
          if (v1 + v2 > v3 && v2 + v3 > v1 && v1 + v3 > v2) {
            var newMax = v1 + v2 + v3;
            if (newMax > max) {
              max = newMax;
            }
          }
        }
      }
    }
    return max;
}

/**
 * Check does array contain triangle.
 *
 * @param {array} arr Input array with numbers.
 * @return {number} 0 in case array doesn't contain triangle, 1 - array contain triangle.
 */
function f3(arr) {
  if (arr.length < 3) return 0;
  arr.sort(function(a, b) { return a - b; });
  for (var i = 0; i < arr.length - 2; i++) {
    var v1 = arr[i];
    var v2 = arr[i + 1];
    var v3 = arr[i + 2];
    if (
      (v1 + v2 > v3 && v2 + v3 > v1 && v1 + v3 > v2)
      && ( (v1 > 0) && (v2 > 0) && (v3 > 0) )
    ) {
      return 1;
    }
  }
  return 0;
}

// console.log(f1([10, 2, 5, 1, 8, 20])); // 23
console.log(f1([1, 1, 1, 3, 3])); //

// console.log(f3([10, 2, 5, 1, 8, 20])); // 1
// console.log(f3([5, 10, 18, 7, 8, 3])); // 1
// console.log(f3([5, 3, 3])); // 1
// console.log(f3([-100, 2, 4, 5])); // 1
// console.log(f3([1, 1, 2, 3, 5])); // 0
// console.log(f3([10, 50, 5, 1])); // 0
// console.log(f3([])); // 0
