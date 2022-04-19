var r = (function(x, f = () => x) {
  var x;
  var y = x;
  x = 2;
  return [x, y, f()];
})(1);
console.log(r); // [ 2, 1, 1 ]

