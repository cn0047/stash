// 1, 1, 2, 3, 5, 8, 13, 21...

// General implementation.
let fibonacci = function (n) {
  if (n < 2) {
    return 1;
  }

  return f(n - 1) + f(n - 2);
};

let f = function (n) {
  if (n === 1 || n === 2) {
    return 1;
  }
  let a = 1;
  let b = 1;
  for (let i = 2; i < n; i++) {
    let v = a + b;
    a = b;
    b = v;
  }
  return b;
};

console.log(f(1) === 1);
console.log(f(2) === 1);
console.log(f(3) === 2);
console.log(f(4) === 3);
console.log(f(5) === 5);
console.log(f(6) === 8);
console.log(f(7) === 13);
console.log(f(8) === 21);
