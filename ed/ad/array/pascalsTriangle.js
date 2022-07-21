// [
//      [1],
//      [1,1],
//      [1,2,1],
//      [1,3,3,1],
//      [1,4,6,4,1]
// ]

function f(n) {
  if (n === 0) return [];
  var arr = [[1]];
  for (var r = 1; r < n; r++) {
    var a = [];
    for (var i = 0; i < r + 1; i++) {
      var v1 = (typeof arr[r - 1][i] === 'undefined') ? 0 : arr[r - 1][i];
      var v2 = (typeof arr[r - 1][i - 1] === 'undefined') ? 0 : arr[r - 1][i - 1];
      a.push(v1 + v2);
    }
    arr.push(a);
  }
  return arr;
}

console.log(f(0));
console.log(f(5));
