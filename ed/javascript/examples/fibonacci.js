// 1, 1, 2, 3, 5, 8, 13...

let f = function (n) {
  if (n === 1 || n === 2) {
    return 1;
  }
  let a = 1;
  let b = 1;
  for (let i = 2; i <= n; i++) {
    let n = a + b;
    a = b;
    b = n;
  }
  return b;
};

console.log(f(2) === 1);
console.log(f(3) === 3);
console.log(f(4) === 5);
console.log(f(5) === 8);
