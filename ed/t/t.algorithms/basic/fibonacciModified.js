// 0, 1, 1, 2, 5, 27...

let f = function (n) {
  if (n === 1) {
    return 0;
  }
  if (n === 2) {
    return 1;
  }
  let a = 0;
  let b = 1;
  for (let i = 3; i <= n; i++) {
    let v = a + (b * b);
    a = b;
    b = v;
  }
  return b;
};

console.log(f(1) == 0);
console.log(f(2) == 1);
console.log(f(3) == 1);
console.log(f(4) == 2);
console.log(f(5) == 5);
console.log(f(6) == 27);
console.log(f(7) == 734);
console.log(f(8) == 538783);
console.log(f(9) == 290287121823);
console.log(f(10) == 84266613096281243382112);
console.log(f(11) == 7100862082718357559748563880517486086728702367);
