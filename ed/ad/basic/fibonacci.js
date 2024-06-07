// 1, 1, 2, 3, 5, 8, 13, 21...

// General implementation.
// @uses: recursion.
const fibonacciV1 = function (n) {
  if (n < 2) {
    return 1;
  }

  return fibonacciV1(n - 1) + fibonacciV1(n - 2);
};

const fibonacciV2 = function (n) {
  if (n === 1 || n === 2) {
    return 1;
  }
  let a = 1;
  let b = 1;
  for (let i = 2; i < n; i++) {
    const v = a + b;
    a = b;
    b = v;
  }
  return b;
};

const f = fibonacciV2;

console.log(f(1) === 1);
console.log(f(2) === 1);
console.log(f(3) === 2);
console.log(f(4) === 3);
console.log(f(5) === 5);
console.log(f(6) === 8);
console.log(f(7) === 13);
console.log(f(8) === 21);
