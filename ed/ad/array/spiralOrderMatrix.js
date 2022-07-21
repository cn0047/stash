function f(a) {
  var r = 0;                   // row
  var c = 0;                   // column
  var rLast = a.length - 1;    // index to last element in row
  var cLast = a[0].length - 1; // index to last element in column

  var result = [];
  while (r <= rLast && c <= cLast) {
    // top line.
    for (var i = c; i <= cLast; i++) {
      result.push(a[r][i]);
      //               ^
    }
    r += 1;
    // right line.
    for (var i = r; i <= rLast; i++) {
      result.push(a[i][cLast]);
      //            ^
    }
    cLast -= 1;
    // bottom line.
    if (r <= rLast) {
      for (var i = cLast; i >= c; i--) {
        result.push(a[rLast][i]);
        //                   ^
      }
      rLast -= 1;
    }
    // left line.
    if (c <= cLast) {
      for (var i = rLast; i >= r; i--) {
        result.push(a[i][c]);
        //            ^
      }
      c += 1;
    }
  }

  return result;
}

var a = [[1]];
var a = [[119]];
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
var a = [
  [ 1,  2,  3,  4],
  [12, 13, 14,  5],
  [11, 16, 15,  6],
  [10,  9,  8,  7],
];
// [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
console.log(f(a));
