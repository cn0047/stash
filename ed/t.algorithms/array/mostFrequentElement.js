/**
 * Get position in array of most frequent element.
 *
 * @param {array} a Array with numbers.
 *
 * @return {number} Index in array of most frequent element.
 */
function f(a) {
  var a = [...arr];
  if (a.length === 0) {
    return -1;
  }
  a.sort(function (a, b) {
    return a - b;
  });
  var frequentElement = a[0];
  var frequentElementCount = 1;
  var currentElement = a[0];
  var currentElementCount = 1;
  for (var i = 1; i < a.length; i++) {
    if (currentElement === a[i]) {
      currentElementCount += 1;
    } else {
      currentElementCount = 1;
    }
    currentElement = a[i];
    if (currentElementCount > frequentElementCount) {
      frequentElement = currentElement;
      frequentElementCount = currentElementCount;
    }
  }
  var isFrequentElementDominator = frequentElementCount > (a.length / 2);
  return isFrequentElementDominator ? arr.indexOf(frequentElement) : -1;
}

// console.log(f([3, 4, 3, 2, 3, -1, 3, 3]));
// console.log(f([]));
// console.log(f([2, 1, 4, 7, 4, 8, 3, 6, 4, 7]));
// console.log(f([0, 0, 1, 1, 1]));
console.log(f([1, 2, 1]));
