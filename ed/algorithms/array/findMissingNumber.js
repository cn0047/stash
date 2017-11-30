/*

OPTION # 1
Missing n = (SUM of numbers from 1 to n) - (SUM all numbers in array)
SUM of numbers from 1 to n = (n * (n + 1)) / 2

OR
OPTION # 2
(XOR range 1 to n) + (XOR actual numbers in array)

*/

/**
 * Finds min missing POSITIVE number in array WITHOUT DUPLICATES.
 *
 * @param {array} A Input array.
 * @return {number} Missing POSITIVE number.
 */
function f(A) {
  // OPTION # 1
  let n = A.length + 1;
  let v1 = (n * (n + 1)) / 2;
  let s = 0;
  for (let i = 0; i < A.length; i++) {
    s += A[i];
  }
  return v1 - s;
}

console.log(f([2, 3, 1, 5]));
