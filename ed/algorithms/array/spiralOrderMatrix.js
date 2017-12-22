function f(a) {
  var r = 0;
  var c = 0;
  var rCount = a.length - 1;
  var cCount = a[0].length - 1;
  var result = [];
  while (r <= rCount && c <= cCount) {
    for (var i = c; i <= cCount; i++) {
      result.push(a[r][i]);
    }
    r += 1;
    for (var i = r; i <= rCount; i++) {
      result.push(a[i][cCount]);
    }
    cCount -= 1;
    if (r <= rCount) {
      for (var i = cCount; i >= c; i--) {
        result.push(a[rCount][i]);
      }
      rCount -= 1;
    }
    if (c <= cCount) {
      for (var i = rCount; i >= r; i--) {
        result.push(a[i][c]);
      }
      c += 1;
    }
  }
  return result;
}

var a = [
  [1, 2, 3],
  [4, 5, 6],
  [7, 8, 9],
];
// [1, 2, 3, 6, 9, 8, 7, 4, 5]
var a = [
  [1, 2, 3],
  [4, 5, 6],
];
// 1 2 3 6 5 4 
var a = [[1]];
var a = [[119]];
console.log(f(a));
