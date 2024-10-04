const add = (n1, n2) => {
  const a = String(n1);
  const b = String(n2);
  const n = a.length > b.length ? a.length : b.length;
  const ad = n - a.length;
  const bd = n - b.length;
  let res = '';
  let addVal = 0;

  for (let i = 0; i < n; i++) {
    let av = parseInt(a[n - ad - i - 1]) || 0;
    let bv = parseInt(b[n - bd - i - 1]) || 0;
    let v = String(av + bv + addVal);
    if (v.length === 2) {
      addVal = parseInt(v[0]);
      v = v[1];
    } else {
      addVal = 0;
    }
    res = String(v) + String(res);
  }
  if (addVal > 0) {
    res = String(addVal) + String(res);
  }

  return res;
}

const multiply = (n1, n2) => {
  const a = String(n1);
  const b = String(n2);
  let sumUp = [];

  for (let bi = b.length - 1; bi >= 0; bi--) {
    let addVal = '';
    // Add zero to the end.
    let r = (new Array(b.length - bi)).join(0);
    for (let ai = a.length - 1; ai >= 0; ai--) {
      let av = parseInt(a[ai]);
      let bv = parseInt(b[bi]);
      let v = String((av * bv) + addVal);
      if (v.length === 2) {
        addVal = parseInt(v[0]);
        v = v[1]
      } else {
        addVal = 0;
      }
      r = String(v) + String(r);
    }
    if (addVal > 0) {
      r = String(addVal) + String(r);
    }
    sumUp.push(r);
  }

  let res = sumUp.reduce(function (sum, v) {
    return add(sum, v);
  }, '0');

  return res;
}

// karatsuba represents algorithm: Karatsuba Multiplication.
// @see: https://en.wikipedia.org/wiki/Karatsuba_algorithm
// @see: https://www.geeksforgeeks.org/karatsuba-algorithm-for-fast-multiplication-using-divide-and-conquer-algorithm
const karatsuba = (num1, num2) => {
  // @TODO.
}

const testAdd = () => {
  console.log(add('0', '001') == 1);
  console.log(add('937', '12188') == (937 + 12188));
  console.log(add('0109', '11122200898') == (109 + 11122200898));
  console.log(add(99, 99) == (99 + 99));
  console.log(add(125, 1952) == (125 + 1952));
  console.log(add(2390525, 3021952)  == (2390525 + 3021952));
  console.log(add(111222333, 999000222111000) == (111222333 + 999000222111000));
}

const testMultiply = () => {
  // const f = karatsuba;
  const f = multiply;
  console.log(f(11, 22) == (11 * 22));
  console.log(f(125, 4282) == (125 * 4282));
  console.log(f(982, 101) == (982 * 101));
  console.log(f(2, 99) == (2 * 99));
  console.log(f(523, 798) == (523 * 798));
  console.log(f('3141592653589793238462643383279502884197169399375105820974944592', '2718281828459045235360287471352662497757247093699959574966967627') == '8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184'); // tested on coursera.
}

(() => {
  testAdd();
  testMultiply();
})();
